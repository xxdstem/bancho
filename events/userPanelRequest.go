package events

import (
	"bancho/common"
	"bancho/packets"
	"bancho/packets/userPackets"
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
		if v == 999 {
			ps.S.Push(packets.BotData())
			continue
		}
		uSession, ok := common.UidToSession[v]
		if !ok {
			continue
		}
		ps.S.Push(userPackets.UserData(uSession.User))
	}
}
