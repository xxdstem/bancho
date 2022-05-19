package handlers

import (
	"bancho/common"
	"bancho/events"
	matches "bancho/events/matches"
	"bancho/packets"
	"bancho/common/log"
)

func HandleEvent(ps common.PackSess) {
	if ps.P.ID == 4 {
		return
	}
	log.Debug("User %d: Requests PacketID: %d", ps.S.User.ID, ps.P.ID)
	switch ps.P.ID {
		case packets.OsuSendIRCMessage:
			events.HandlePublicMessage(ps)
		case packets.OsuSendIRCMessagePrivate:
			events.HandlePrivateMessage(ps)
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
		case packets.OsuExit:
			events.LogOut(ps)
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
