package events

import (
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/common"
	
)

// LoginData is the data received by the osu! client upon a login request to bancho.

func UpdateStats(ps common.PackSess){
	var newMode byte
	ps.P.Unmarshal(
		&ps.S.User.Status.Status,
		&ps.S.User.Status.Text,
		&ps.S.User.Status.MD5,
		&ps.S.User.Status.Mods,
		&newMode,
		&ps.S.User.Status.BeatmapID,
	)
	if ps.S.User.Stats.Mode  != newMode{
		ps.S.User.UpdateStats(newMode)
	}
	ps.S.Push(packets.UserDataFull(&ps.S.User))
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

