package events

import (
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/common"
	
)

func UserStatsRequest(ps common.PackSess) {
	var usersRequested []int32
	err := ps.P.Unmarshal(&usersRequested)
	if err != nil {
		return
	}
	common.UidToSessionMutex.Lock()
	
	common.UidToSessionMutex.Unlock()
	for _, v := range usersRequested {
		uSession, ok := common.UidToSession[v]
		if !ok {
			continue
		}
		ps.S.Push(packets.UserData(&uSession.User))
	}
}