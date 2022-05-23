package events

import (
	"bancho/common"
	"bancho/packets"
)

func ChangeMods(ps common.PackSess) {
	match := ps.S.User.Match
	mods := packets.MatchMods(ps.P)
	match.Mutex.Lock()
	defer match.Mutex.Unlock()
	match.Settings.Mods = mods
	UpdateMatch(match)
}
