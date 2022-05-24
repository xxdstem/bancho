package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
)

func ChangeSlot(ps common.PackSess) {
	match := ps.S.User.Match
	slotID := packets.Slot(ps.P)
	log.Debug("User %d: Changed slot to %d", ps.S.User.ID, slotID)
	match.Mutex.Lock()
	defer match.Mutex.Unlock()
	for id, slot := range match.Players {
		if slot.User != nil && slot.User.ID == ps.S.User.ID {
			match.Players[slotID] = match.Players[id]
			match.Players[id].Status = 1
			match.Players[id].User = nil
			UpdateMatch(match)
			return
		}
	}

}
