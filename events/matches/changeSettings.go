package events

import (
	"bancho/common"
	"bancho/packets"
)

func ChangeSettings(ps common.PackSess) {
	match := ps.S.User.Match
	packetData := packets.MatchSettings(ps.P)
	match.Mutex.Lock()
	defer match.Mutex.Unlock()
	match.Beatmap = common.MatchBeatmap{
		Name: packetData.BeatmapName,
		MD5:  packetData.BeatmapMD5,
		ID:   packetData.BeatmapID,
	}
	UpdateMatch(ps)
}
