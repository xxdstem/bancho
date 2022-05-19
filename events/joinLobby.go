package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/common/log"
)

func JoinLobby(ps common.PackSess) {
	s := common.GetStream("lobby")
	s.Subscribe(ps.S.User.Token)
	common.MatchesMutex.Lock()
	defer common.MatchesMutex.Unlock()
	for _, v := range common.Matches {
		if v != nil {
			ps.S.Push(packets.MatchDataFull(v, packets.BanchoMatchNew, true))
		}
	}
	log.Debug("User %d: Joined lobby", ps.S.User.ID)
}
