package events

import (
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/common"
	"fmt"	
)

func UserPanelRequest(ps common.PackSess) {
	var usersRequested []int32
	err := ps.P.Unmarshal(&usersRequested)
	if err != nil {
		return
	}
	common.UidToSessionMutex.Lock()
	defer common.UidToSessionMutex.Unlock()
	for _, v := range usersRequested {
		if v == 999{
			ps.S.Push(packets.BotData())
			continue
		}
		uSession, ok := common.UidToSession[v]
		if !ok {
			continue
		}
		ps.S.Push(packets.UserData(&uSession.User))
	}
}