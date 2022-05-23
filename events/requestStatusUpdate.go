package events

import (
	"bancho/common"
	"bancho/packets"
)

func ReturnUserStats(ps common.PackSess) {
	ps.S.Push(packets.UserDataFull(ps.S.User))
}
