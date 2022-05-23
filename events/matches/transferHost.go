package events

import (
	"bancho/common"
	"bancho/packets"
)

func TransferHost(ps common.PackSess) {
	slotID := packets.Slot(ps.P)
	match := ps.S.User.Match
	match.Mutex.Lock()
	match.TransferHost(slotID)
	UpdateMatch(match)
	match.Mutex.Unlock()
}
