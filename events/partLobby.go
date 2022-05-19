package events

import (
	"bancho/common"
	"bancho/common/log"
)

func PartLobby(ps common.PackSess) {
	s := common.GetStream("lobby")
	s.Unsubscribe(ps.S.User.Token)
	log.Debug("User %d: Parted lobby", ps.S.User.ID)
}
