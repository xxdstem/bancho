package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
)

func UpdateMatch(ps common.PackSess){
	match := ps.S.User.Match
	if match != nil{
		ps.S.Push(packets.MatchDataFull(match, packets.BanchoMatchUpdate, false))
	}
	
}