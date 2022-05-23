package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
)

func PartMatch(ps common.PackSess) {
	if ps.S.User.Match == nil {
		log.Error("some error. i got NULL match?")
		return
	}
	match := ps.S.User.Match
	matchID := ps.S.User.Match.ID
	ps.S.Push(packets.ChannelKicked(ps.S.User.Match.Channel.ClientName))

	dispose := ps.S.User.LeaveMatch()

	if dispose {
		s := common.GetStream("lobby")
		pack := packets.DisposeMatch(matchID)
		s.Send(pack)
		ps.S.Push(pack)
		log.Debug("User %d: Disposing match #%d", ps.S.User.ID, matchID)
	} else {
		log.Debug("User %d: Updating match #%d", ps.S.User.ID, matchID)
		UpdateMatch(match)
	}
	log.Debug("User %d: Leaving match #%d", ps.S.User.ID, matchID)
}
