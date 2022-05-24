package events

import (
	"bancho/common"
	"bancho/packets"
)

func ChangeSettings(ps common.PackSess) {
	match := ps.S.User.Match
	packetData := packets.MatchSettings(ps.P)
	match.Mutex.Lock()
	match.Beatmap = common.MatchBeatmap{
		Name: packetData.BeatmapName,
		MD5:  packetData.BeatmapMD5,
		ID:   packetData.BeatmapID,
	}
	mods := match.Settings.Mods
	if packetData.ModMode != match.Settings.ModMode {
		mods = 0
	}
	match.Settings = common.MatchSettings{
		GameMode:    packetData.GameMode,
		Mods:        mods,
		ScoringType: packetData.ScoringType,
		TeamType:    packetData.TeamType,
		ModMode:     packetData.ModMode,
	}

	UpdateMatch(match)
	match.Mutex.Unlock()
}
