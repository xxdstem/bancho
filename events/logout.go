package events

import (
	"bancho/common"
	"bancho/packets"
)

func LogOut(ps common.PackSess) {
	s := common.GetStream("main")
	s.Send(packets.LogOut(ps.S.User.ID))
	common.DeleteSession(ps.S)
}
