package common

import (
	"fmt"
)

const statsQuery = `
SELECT pp_?, playcount_? FROM user_stats WHERE id = ?
`

// User represents an user online on bancho.
type User struct {
	ID        	int32
	Name      	string
	Token     	string
	UTCOffset 	byte
	Country   	byte
	Colour    	byte
	Stats		UserStats
	Status		UserStatus
	Position  	struct {
		Longitude float32
		Latitude  float32
	}
}

type UserStatus struct{
	Status    byte
	Text      string
	MD5    string
	Mods      int32
	BeatmapID int32
}

type UserStats struct{
	PP		 	uint16
	Rank	 	uint32
	PlayCount	uint32
	GameMode 	byte
}

func (u *User) UpdateStats(mode int) {
	modeText := IntToGameMode(mode)
	err := DB.QueryRow(statsQuery, modeText, modeText, u.ID).Scan(&u.Stats.PP, &u.Stats.PlayCount)
	if err != nil{
		fmt.Println(err)
	}
	// do updates/
}

func IntToGameMode(mode int) string{
	switch (mode){
		default:
			return "std"
		case 1:
			return "taiko"
		case 2:
			return "ctb"
		case 3:
			return "mania"
	}
}