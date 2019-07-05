package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"fmt"
	"sync"
)

func CreateMatch(ps common.PackSess){
	packetData := packets.MatchSettings(ps.P)
	b := common.MatchBeatmap{
		Name:	packetData.BeatmapName,
		MD5:	packetData.BeatmapMD5,
		ID:		packetData.BeatmapID,
	}
	m := common.Match{
		Name:			packetData.Name,
		CreatorID: 		packetData.CreatorID,
		HostID: 		packetData.CreatorID,
		Mutex: 			&sync.Mutex{},
		Beatmap:		b,
	}
	for i := 0; i < packetData.Slots; i++ {
		m.Players[i].Status = 1
	}
	for i := packetData.Slots; i < 16; i++ {
		m.Players[i].Status = 2
	}
	match := common.NewMatch(m)
	fmt.Println("created match #", match.ID)
	ps.S.User.JoinMatch(match)
	s := common.GetStream("lobby")
	s.Send(packets.MatchDataFull(match, packets.BanchoMatchNew, true))
	ps.S.Push(packets.MatchDataFull(match, packets.BanchoMatchJoinSuccess, false))
}