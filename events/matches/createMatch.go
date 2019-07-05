package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"fmt"
	"sync"
)

func CreateMatch(ps common.PackSess){
	var match  *common.Match
	packetData := packets.MatchSettings(ps.P)
	match = &common.Match{
		Name:			packetData.Name,
		CreatorID: 		packetData.CreatorID,
		HostID: 		packetData.CreatorID,
		Mutex: 			&sync.Mutex{},
	}
	match.Beatmap = common.MatchBeatmap{
		Name:	packetData.BeatmapName,
		MD5:	packetData.BeatmapMD5,
		ID:		packetData.BeatmapID,
	}
	fmt.Println(match)
	for i := 0; i < packetData.Slots; i++ {
		match.Players[i].Status = 1
	}
	for i := packetData.Slots; i < 16; i++ {
		match.Players[i].Status = 2
	}
	match.Players[0].User = &ps.S.User
	match.Players[0].Team = 0
	match.Players[0].Status = 4
	ps.S.User.Match = common.NewMatch(match)
	fmt.Println("created match #", ps.S.User.Match.ID)
	ps.S.Push(packets.MatchDataFull(match, packets.BanchoMatchJoinSuccess))
}