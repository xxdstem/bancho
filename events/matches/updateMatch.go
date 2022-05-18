package events

import (
	"bancho/common"
	"bancho/packets"
)

func UpdateMatch(ps common.PackSess) {
	match := ps.S.User.Match
	if match != nil {
		ps.S.Push(packets.MatchDataFull(match, packets.BanchoMatchUpdate, false))
	}

}
