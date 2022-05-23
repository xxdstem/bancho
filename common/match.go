package common

import (
	"bancho/common/log"
	"sync"
)

type Match struct {
	ID         uint32
	Name       string
	Password   string
	CreatorID  int32
	HostID     int32
	InProgress bool
	Settings   MatchSettings
	Beatmap    MatchBeatmap
	Players    [16]MatchPlayer
	Mutex      *sync.Mutex
}

type MatchSettings struct {
	GameMode    int
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
	m.ID = uint32(lastMatchID)
	Matches[lastMatchID] = &m

	return Matches[lastMatchID]

}

func DisposeMatch(m *Match) {
	MatchesMutex.Lock()
	defer MatchesMutex.Unlock()
	Matches[int(m.ID)] = nil
	log.Debug("disposing match %d", m.ID)
}

func (m *Match) getUserSlotID(u int32) int {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	for i, slot := range m.Players {
		if slot.User.ID == u {
			return i
		}
	}
	return -1
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
	slotID := m.getUserSlotID(u.ID)
	m.Players[slotID].User = nil
	m.Players[slotID].Team = 0
	m.Players[slotID].Status = 1
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if m.countUsers() == 0 {
		DisposeMatch(m)
		return true, true
	}
	return true, false

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
