package packets

import (
	//"github.com/xxdstem/bancho/common"
	"bancho/inbound"
	"math"
)

type SettignsStruct struct {
	MatchID       uint16
	InProgress    byte
	unknown       byte
	Mods          uint32
	Name          string
	Password      string
	BeatmapName   string
	BeatmapID     uint32
	BeatmapMD5    string
	slotStatus    [16]byte
	slotTeams     [16]byte
	CreatorID     int32
	PlayMode      byte
	ScoreingType  byte
	MatchTeamType byte
	Slots         int
}

func JoinMatch(pack inbound.BasePacket) (uint32, string) {
	var (
		id uint32
		pw string
	)
	pack.Unmarshal(&id, &pw)
	return id, pw
}

func Slot(pack inbound.BasePacket) uint32 {
	var slotID uint32
	pack.Unmarshal(&slotID)
	return slotID
}

func MatchMods(pack inbound.BasePacket) int32 {
	var m int32
	pack.Unmarshal(&m)
	return m
}

func MatchSettings(pack inbound.BasePacket) SettignsStruct {
	m := SettignsStruct{}
	pack.Unmarshal(
		&m.MatchID,
		&m.InProgress,
		&m.unknown,
		&m.Mods,
		&m.Name,
		&m.Password,
		&m.BeatmapName,
		&m.BeatmapID,
		&m.BeatmapMD5,
		&m.slotStatus,
		&m.slotTeams,
		&m.CreatorID,
		&m.PlayMode,
		&m.ScoreingType,
		&m.MatchTeamType,
	)
	var i int
	for _, v := range m.slotStatus {
		if v == 1 {
			i++
		}
	}
	m.Slots = int(math.Min(float64(i), 16))
	return m
}
