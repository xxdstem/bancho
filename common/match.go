package common


type Match struct {
	ID        	int32
	Name      	string
	CreatorID	int32
	Players		[]MatchPlayer
}

type MatchPlayer struct{
	User	User
	Score	PlayerScore
	Mods	int32
}

type PlayerScore struct{
	Count300	int32
	Count100	int32
	Count50		int32
	CountMiss	int32
	Score		uint64
	Combo		int32
}


func (u *MatchPlayer) UpdateScore(score uint64) {
	u.Score.Score = score
}
