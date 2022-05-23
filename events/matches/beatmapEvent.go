package events

import (
	"bancho/common"
)

func BeatmapEvent(ps common.PackSess, has bool) {
	match := ps.S.User.Match
	match.UserBeatmapStatus(ps.S.User, has)
	UpdateMatch(match)
}
