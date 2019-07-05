package events

import (
	"github.com/xxdstem/bancho/common"
	"fmt"
)

func CreateMatch(ps common.PackSess){
	//var match  *common.Match
	//match = &common.Match{}
	var a uint16
	var prog, u, playmode, scoringtype, matchteamtype byte
	var mods uint32
	var name, password, bname, bmd5 string
	var bid uint32
	var slotStatus [16]byte
	var slotTeams [16]byte
	var hostID int32
	ps.P.Unmarshal(
		&a,
		&prog,
		&u,
		&mods,
		&name,
		&password,
		&bname,
		&bid,
		&bmd5,
		&slotStatus,
		&slotTeams,
		&hostID,
		&playmode,
		&scoringtype,
		&matchteamtype,
	)
	fmt.Println(name, password, slotStatus, hostID, playmode)
}