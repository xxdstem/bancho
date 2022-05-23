package events

import (
	"bancho/common"
	"bancho/packets/userPackets"
)

func ReturnUserStats(ps common.PackSess) {
	ps.S.Push(userPackets.UserDataFull(ps.S.User))
}
