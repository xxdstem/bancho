package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
)

func PartLobby(ps common.PackSess) {
	s := common.GetStream("lobby")
	s.Unsubscribe(ps.S.User.Token)
	ch := chat.GetChannel("#lobby")
	ps.S.User.LeaveChannel(ch)
	log.Debug("User %d: Parted lobby", ps.S.User.ID)
}
