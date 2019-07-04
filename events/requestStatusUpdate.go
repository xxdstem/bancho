package events

import (
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/common"
	
)

// LoginData is the data received by the osu! client upon a login request to bancho.

func ReturnUserStats(ps common.PackSess){
	ps.S.Push(packets.UserDataFull(&ps.S.User))
}

