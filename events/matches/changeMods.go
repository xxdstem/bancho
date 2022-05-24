package events

import (
	"bancho/common"
	"bancho/packets"
)

func ChangeMods(ps common.PackSess) {
	match := ps.S.User.Match
	mods := packets.MatchMods(ps.P)
	match.Mutex.Lock()
	if match.Settings.ModMode == 0 {
		match.Settings.Mods = mods
	} else { // FreeMod
		if match.HostID == ps.S.User.ID {
			//Set DT or NC global.
			if (mods & 64) > 0 {
				match.SetMods(64)
				if (mods & 512) > 0 {
					match.SetMods(512 + 64)
				}
			} else if (mods & 256) > 0 {
				match.SetMods(256)
			} else {
				match.SetMods(0)
			}
		}
		match.SetPlayerMods(ps.S.User.ID, mods)
	}

	UpdateMatch(match)
	match.Mutex.Unlock()
}
