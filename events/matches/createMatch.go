package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
	"sync"
)

func CreateMatch(ps common.PackSess) {
	packetData := packets.MatchSettings(ps.P)
	b := common.MatchBeatmap{
		Name: packetData.BeatmapName,
		MD5:  packetData.BeatmapMD5,
		ID:   packetData.BeatmapID,
	}
	m := common.Match{
		Name:      packetData.Name,
		CreatorID: packetData.CreatorID,
		HostID:    packetData.CreatorID,
		Mutex:     &sync.Mutex{},
		Beatmap:   b,
	}
	for i := 0; i < packetData.Slots; i++ {
		m.Players[i].Status = 1
	}
	for i := packetData.Slots; i < 16; i++ {
		m.Players[i].Status = 2
	}
	match := common.NewMatch(m)
	log.Debug("User %d: Created match #%d", ps.S.User.ID, match.ID)
	ps.S.User.JoinMatch(match)
	s := common.GetStream("lobby")
	s.Send(userPackets.MatchDataFull(match, packets.BanchoMatchNew, true))
	ps.S.Push(userPackets.MatchDataFull(match, packets.BanchoMatchJoinSuccess, false))
}
