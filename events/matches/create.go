package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
	"fmt"
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
		CreatorID: ps.S.User.ID,
		HostID:    ps.S.User.ID,
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
	match.Channel = chat.NewChannel(fmt.Sprintf("#multi_%d", match.ID), true)

	log.Debug("User %d: Created match #%d", ps.S.User.ID, match.ID)
	ps.S.User.JoinMatch(match)
	ps.S.User.JoinChannel(match.Channel)
	ps.S.Push(userPackets.ChannelJoin(match.Channel))
	s := common.GetStream("lobby")
	s.Send(userPackets.MatchDataFull(match, packets.BanchoMatchNew, true))
	ps.S.Push(userPackets.MatchDataFull(match, packets.BanchoMatchJoinSuccess, false))
}
