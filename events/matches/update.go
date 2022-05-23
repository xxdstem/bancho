package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/packets/userPackets"
)

func UpdateMatch(match *common.Match) {
	if match != nil {
		match.Stream.Send(userPackets.MatchDataFull(match, packets.BanchoMatchUpdate, false))
	}

}
