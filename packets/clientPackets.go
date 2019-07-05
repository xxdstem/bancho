package packets

import (
	//"github.com/xxdstem/bancho/common"
	"git.zxq.co/ripple/nuclearbancho/inbound"
	"math"
)

type SettinsStruct struct{
	MatchID			uint16
	InProgress		byte
	unknown			byte
	Mods			uint32
	Name			string
	Password		string
	BeatmapName		string
	BeatmapID		uint32
	BeatmapMD5		string
	slotStatus  	[16]byte
	slotTeams   	[16]byte
	CreatorID		int32
	PlayMode		byte
	ScoreingType	byte
	MatchTeamType	byte
	Slots			int
}

func MatchSettings(pack inbound.BasePacket) SettinsStruct{
	m := SettinsStruct{}
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
	for _, v := range m.slotStatus{
		if v == 1{
			i++
		}
	}
	m.Slots = int(math.Min(float64(i), 16))
	return m
}