package handlers

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/events"
	matches "github.com/xxdstem/bancho/events/matches"
	"fmt"

)

func HandleEvent(ps common.PackSess){
	if ps.P.ID == 4{
		return
	}
	fmt.Println("Requested packID: ",  ps.P.ID)
	switch ps.P.ID{
		case packets.OsuSendUserState:
			events.UpdateStats(ps)
		case packets.OsuUserStatsRequest:
			events.UserStatsRequest(ps)
		case packets.OsuUserPresenceRequest:
			events.UserPanelRequest(ps)
		case packets.OsuRequestStatusUpdate:
			events.ReturnUserStats(ps)
		case packets.OsuLobbyJoin:
			events.JoinLobby(ps)
		case packets.OsuLobbyPart:
			events.PartLobby(ps)
		case packets.OsuMatchCreate:
			matches.CreateMatch(ps)
		case packets.OsuMatchChangeSlot:
			matches.ChangeSlot(ps)
		case packets.OsuMatchChangeSettings:
			matches.ChangeSettings(ps)
		case packets.OsuMatchPart:
			matches.PartMatch(ps)
	}
}


