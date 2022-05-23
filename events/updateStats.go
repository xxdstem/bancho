package events

import (
	"bancho/common"
	"bancho/packets/userPackets"
)

func UpdateStats(ps common.PackSess) {
	var newMode byte
	ps.P.Unmarshal(
		&ps.S.User.Status.Status,
		&ps.S.User.Status.Text,
		&ps.S.User.Status.MD5,
		&ps.S.User.Status.Mods,
		&newMode,
		&ps.S.User.Status.BeatmapID,
	)
	if ps.S.User.Stats.Mode != newMode {
		ps.S.User.UpdateStats(newMode)
	}
	ps.S.Push(userPackets.UserDataFull(ps.S.User))
	/*
		var usersRequested []int32
		err := ps.p.Unmarshal(&usersRequested)
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
			ps.s.Push(packets.UserData(&uSession.User))
		}*/
}
