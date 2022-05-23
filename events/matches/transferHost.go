package events

import (
	"bancho/common"
	"bancho/packets"
)

func TransferHost(ps common.PackSess) {
	slotID := packets.Slot(ps.P)
	match := ps.S.User.Match
	match.Mutex.Lock()
	defer match.Mutex.Unlock()
	match.TransferHost(slotID)
	UpdateMatch(match)
}
