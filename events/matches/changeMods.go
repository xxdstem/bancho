package events

import (
	"bancho/common"
	"bancho/packets"
)

func ChangeMods(ps common.PackSess) {
	match := ps.S.User.Match
	mods := packets.MatchMods(ps.P)
	match.Mutex.Lock()
	match.Settings.Mods = mods
	UpdateMatch(match)
	match.Mutex.Unlock()
}
