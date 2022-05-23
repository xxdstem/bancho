package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
)

func JoinMatch(ps common.PackSess) {
	matchID, _ := packets.JoinMatch(ps.P)
	m := common.GetMatch(matchID)
	ok := ps.S.User.JoinMatch(m)
	if !ok {
		ps.S.Push(packets.MatchJoinFailed())
		return
	}

	ps.S.User.JoinChannel(m.Channel)

	ps.S.Push(userPackets.MatchDataFull(m, packets.BanchoMatchJoinSuccess, false))
	ps.S.Push(userPackets.ChannelJoin(m.Channel))

	log.Info("%s joined match #%d", ps.S.User.Name, ps.S.User.Match.ID)
	UpdateMatch(m)

}
