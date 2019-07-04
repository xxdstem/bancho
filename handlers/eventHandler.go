package handlers

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/events"
	"fmt"

)

func HandleEvent(ps common.PackSess){
	fmt.Println(ps.P.ID)
	switch ps.P.ID{
		case packets.OsuSendUserState:
			events.UpdateStats(ps)
		case packets.OsuUserStatsRequest:
			events.UserStatsRequest(ps)
		case packets.OsuRequestStatusUpdate:
			events.ReturnUserStats(ps)
	}
}


