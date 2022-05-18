package events

import (
	"bancho/packets"
	"bancho/common"
)

func UserStatsRequest(ps common.PackSess) {
	var usersRequested []int32
	err := ps.P.Unmarshal(&usersRequested)
	if err != nil {
		return
	}
	common.UidToSessionMutex.Lock()
	defer common.UidToSessionMutex.Unlock()
	for _, v := range usersRequested {
		if v == ps.S.User.ID{
			continue
		}
		uSession, ok := common.UidToSession[v]
		if !ok {
			continue
		}
		ps.S.Push(packets.UserDataFull(&uSession.User))
	}
}