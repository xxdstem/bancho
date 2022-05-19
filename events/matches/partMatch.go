package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/common/log"
)

func PartMatch(ps common.PackSess) {
	matchID := ps.S.User.Match.ID
	dispose := ps.S.User.LeaveMatch()
	if dispose {
		s := common.GetStream("lobby")
		s.Send(packets.DisposeMatch(matchID))
		log.Debug("User %d: Disposing match #%d", ps.S.User.ID, matchID)
	} else {
		log.Debug("User %d: Updating match #%d", ps.S.User.ID, matchID)
		UpdateMatch(ps)
	}

	log.Debug("User %d: Leaving match #%d", ps.S.User.ID, matchID)
}
