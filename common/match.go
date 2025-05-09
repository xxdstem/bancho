package common

import (
	"bancho/common/log"
	"fmt"
	"sync"
)

type Match struct {
	ID         uint32
	Name       string
	Password   string
	CreatorID  int32
	HostID     int32
	InProgress bool
	Stream     *Stream
	Channel    *Channel
	Settings   MatchSettings
	Beatmap    MatchBeatmap
	Players    [16]MatchPlayer
	Mutex      *sync.Mutex
}

type MatchSettings struct {
	GameMode    byte
	Mods        int32
	ScoringType byte
	TeamType    byte
	ModMode     byte
}

type MatchBeatmap struct {
	Name string
	MD5  string
	ID   uint32
}

type MatchPlayer struct {
	User   *User
	Score  PlayerScore
	Team   byte
	Status byte
	Failed bool
	Passed bool
	Mods   int32
}

type PlayerScore struct {
	Count300  int32
	Count100  int32
	Count50   int32
	CountMiss int32
	Score     uint64
	Combo     int32
}

func NewMatch(m Match) *Match {
	MatchesMutex.Lock()
	defer MatchesMutex.Unlock()
	lastMatchID++
	m.Stream = NewStream(fmt.Sprintf("multi/%d", lastMatchID))
	m.ID = uint32(lastMatchID)
	Matches[lastMatchID] = &m

	return Matches[lastMatchID]
}

func GetMatch(id uint32) *Match {
	return Matches[int(id)]
}

func DisposeMatch(m *Match) {
	MatchesMutex.Lock()
	defer MatchesMutex.Unlock()
	m.Stream.Delete()
	m.Channel.Stream.Delete()
	Matches[int(m.ID)] = nil
	log.Debug("disposing match %d", m.ID)
}

func (m *Match) getUserSlotID(u int32) *MatchPlayer {
	for i, slot := range m.Players {
		if slot.User != nil && slot.User.ID == u {
			return &m.Players[i]
		}
	}
	return nil
}

func (m *Match) SetPlayerMods(userID int32, mods int32) {
	p := m.getUserSlotID(userID)
	p.Mods = mods & ^(64 | 256 | 512)
	log.Info("setting to %s mods: %d | %d", p.User.Name, p.Mods, mods)
}

func (m *Match) SetMods(mods int32) {
	m.Settings.Mods = mods
}

func (m *Match) ToggleSlotLocked(slotID uint32) bool {
	slot := &m.Players[slotID]
	if slot.User != nil {
		m.UserLeft(slot.User)
		return true
	}
	if slot.Status == 2 {
		slot.Status = 1
	} else {
		slot.Status = 2
	}
	return false
}
func (m *Match) TransferHost(slotID uint32) {
	u := m.Players[slotID].User
	if u != nil {
		m.HostID = u.ID
	}
}

func (m *Match) ToggleReady(userID int32) {
	p := m.getUserSlotID(userID)
	switch p.Status {
	case 4:
		p.Status = 8
	case 8:
		p.Status = 4
	default:
		break
	}
}

func (m *Match) UserBeatmapStatus(u *User, has bool) {
	p := m.getUserSlotID(u.ID)
	if has {
		p.Status = 4
	} else {
		p.Status = 16
	}
}

func (m *Match) UserJoin(u *User) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	var spareSlot int
	spareSlot = -1
	for i, slot := range m.Players {
		if spareSlot == -1 && slot.Status == 1 {
			spareSlot = i
		}
	}
	if spareSlot == -1 {
		return false
	}
	m.Players[spareSlot].User = u
	m.Players[spareSlot].Team = 0
	m.Players[spareSlot].Status = 4
	return true
}

func (m *Match) UserLeft(u *User) (bool, bool) {
	p := m.getUserSlotID(u.ID)
	if p != nil {
		p.User = nil
		p.Team = 0
		p.Status = 1
		if m.countUsers() == 0 {
			DisposeMatch(m)
			return true, true
		}
		return true, false
	}
	return false, false

}

func (m *Match) countUsers() int {
	var i int
	for _, v := range m.Players {
		if v.User != nil {
			i++
		}
	}
	return i
}

func (u *MatchPlayer) UpdateScore(score uint64) {
	u.Score.Score = score
}
