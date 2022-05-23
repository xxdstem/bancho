package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
)

func JoinLobby(ps common.PackSess) {
	s := common.GetStream("lobby")
	s.Subscribe(ps.S.User.Token)
	ch := chat.GetChannel("#lobby")
	ps.S.Push(userPackets.ChannelJoin(ch))
	common.MatchesMutex.Lock()
	defer common.MatchesMutex.Unlock()
	for _, v := range common.Matches {
		if v != nil {
			ps.S.Push(userPackets.MatchDataFull(v, packets.BanchoMatchNew, true))
		}
	}
	log.Debug("User %d: Joined lobby", ps.S.User.ID)
}
