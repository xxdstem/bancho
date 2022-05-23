package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
)

func PartLobby(ps common.PackSess) {
	s := common.GetStream("lobby")
	s.Unsubscribe(ps.S.User.Token)
	ch := chat.GetChannel("#lobby")
	ps.S.User.LeaveChannel(ch)
	ps.S.Push(packets.ChannelKicked("#lobby"))
	log.Debug("User %d: Parted lobby", ps.S.User.ID)
}
