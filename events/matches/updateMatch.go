package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/packets/userPackets"
)

func UpdateMatch(ps common.PackSess) {
	match := ps.S.User.Match
	if match != nil {
		ps.S.Push(userPackets.MatchDataFull(match, packets.BanchoMatchUpdate, false))
	}

}
