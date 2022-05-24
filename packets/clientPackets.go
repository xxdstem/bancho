package packets

import (
	//"github.com/xxdstem/bancho/common"
	"bancho/inbound"
	"fmt"
	"math"
)

type SettignsStruct struct {
	MatchID     uint16
	InProgress  byte
	unknown     byte
	Mods        uint32
	Name        string
	Password    string
	BeatmapName string
	BeatmapID   uint32
	BeatmapMD5  string
	slotStatus  [16]byte
	slotTeams   [16]byte
	GameMode    byte
	ScoringType byte
	TeamType    byte
	ModMode     byte
	Slots       int
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
	)

	start := 7 + 2 + 1 + 1 + 4 + 4 + 16 + 16 + len(m.Name) + len(m.Password) + len(m.BeatmapMD5) + len(m.BeatmapName) + 8
	for _, v := range m.slotStatus {
		if v > 4 {
			start += 4
		}
	}
	if m.Name != "" {
		start += 1
	}
	if m.Password != "" {
		start += 1
	}
	// fix := false
	// l := len(pack.Content)
	// for !fix && l-start > 10 {
	// 	var v byte
	// 	inbound.BasePacket{
	// 		Content: pack.Content[start:],
	// 	}.Unmarshal(&v)
	// 	if v == 0 {
	// 		start += 1
	// 		fmt.Println("havayu")
	// 	} else {
	// 		fix = true
	// 	}
	// }
	pack = inbound.BasePacket{
		Content: pack.Content[start:],
	}
	fmt.Println(pack.Content)
	pack.Unmarshal(
		&m.GameMode,
		&m.ScoringType,
		&m.TeamType,
		&m.ModMode)
	var i int
	for _, v := range m.slotStatus {
		if v == 1 {
			i++
		}
	}
	fmt.Println(m)
	m.Slots = int(math.Min(float64(i), 16))
	return m
}
