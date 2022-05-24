package events

import (
	"bancho/common"
)

func SetReady(ps common.PackSess) {
	match := ps.S.User.Match
	match.Mutex.Lock()
	match.ToggleReady(ps.S.User.ID)
	UpdateMatch(match)
	match.Mutex.Unlock()
}
