package handlers

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/events"

)

func HandleEvent(ps common.PackSess){
	switch ps.P.ID{
		case packets.OsuSendUserState:
			events.UpdateStats(ps)
		case packets.OsuUserStatsRequest:
			events.UserStatsRequest(ps)
	}
}


