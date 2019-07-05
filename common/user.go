package common

import (
	"fmt"
)

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
	Match		*Match
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
	Accuracy	float64
	TotalScore	uint64
	RankedScore uint64
	Mode 		byte
}

func (u *User) JoinMatch(m *Match){
	m.UserJoin(u)
	u.Match = m
}

func (u *User) LeaveMatch() bool{
	_, dispose := u.Match.UserLeft(u)
	u.Match = nil
	return dispose
}

func (u *User) UpdateStats(mode byte) {
	modeText := IntToGameMode(mode)
	statsQuery := `
	SELECT pp_`+modeText+`, playcount_`+modeText+`, avg_accuracy_`+modeText+`/100, ranked_score_`+modeText+`, total_score_`+modeText+`, 0 FROM users_stats WHERE id = ?
	`
	err := DB.QueryRow(statsQuery, u.ID).Scan(&u.Stats.PP, &u.Stats.PlayCount, &u.Stats.Accuracy, &u.Stats.RankedScore, &u.Stats.TotalScore, &u.Stats.Rank)
	if err != nil{
		fmt.Println(err)
	}
	u.Stats.Mode = mode
	// do updates/
}

func IntToGameMode(mode byte) string{
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