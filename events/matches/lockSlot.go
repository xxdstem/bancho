package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
)

func LockSlot(ps common.PackSess) {
	match := ps.S.User.Match
	slotID := packets.Slot(ps.P)
	log.Debug("User %d: Changed slot to %d", ps.S.User.ID, slotID)
	u := match.Players[slotID].User
	if u != nil && u.ID == ps.S.User.ID {
		return
	}
	if match.ToggleSlotLocked(slotID) {
		sess := common.GetSession(u.Token)
		sess.Push(userPackets.MatchDataFull(match, packets.BanchoMatchUpdate, false))
	}
	UpdateMatch(match)

}
