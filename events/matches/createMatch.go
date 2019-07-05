package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"fmt"
	"sync"
)

func CreateMatch(ps common.PackSess){
	var match  *common.Match
	match = &common.Match{Mutex: &sync.Mutex{}}
	var a uint16
	var prog, u, playmode, scoringtype, matchteamtype byte
	var mods uint32
	var password string
	var slotStatus [16]byte
	var slotTeams [16]byte
	ps.P.Unmarshal(
		&a,
		&prog,
		&u,
		&mods,
		&match.Name,
		&password,
		&match.Beatmap.Name,
		&match.Beatmap.ID,
		&match.Beatmap.MD5,
		&slotStatus,
		&slotTeams,
		&match.CreatorID,
		&playmode,
		&scoringtype,
		&matchteamtype,
	)
	match.HostID = match.CreatorID

	for i, _ := range match.Players{ 
		match.Players[i].Team = 0
		match.Players[i].Status = 1
	}
	match.Players[0].User = &ps.S.User
	match.Players[0].Team = 0
	match.Players[0].Status = 4
	ps.S.User.Match = common.NewMatch(match)
	fmt.Println("created match #", ps.S.User.Match.ID)
	ps.S.Push(packets.MatchDataFull(match, packets.BanchoMatchJoinSuccess))
}