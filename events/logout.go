package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/common/log"
)

func LogOut(ps common.PackSess) {
	s := common.GetStream("main")
	s.Send(packets.LogOut(ps.S.User.ID))
	common.DeleteSession(ps.S)
	log.Debug("User %d: Logout.", ps.S.User.ID)
}
