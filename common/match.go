package common

import "sync"

type Match struct {
	ID        	int
	Name      	string
	Password	string
	CreatorID	int32
	HostID		int32
	Beatmap		MatchBeatmap
	Players		[16]MatchPlayer
	Mutex		*sync.Mutex
}

type MatchBeatmap struct{
	Name	string
	MD5		string
	ID		uint32
}

type MatchPlayer struct{
	User	*User
	Score	PlayerScore
	Team	byte
	Status	byte
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

func NewMatch(m *Match) *Match {
	MatchesMutex.Lock()
	lastMatchID++
	m.ID = lastMatchID
	Matches[lastMatchID] = m
	MatchesMutex.Unlock()
	return Matches[lastMatchID]
}

func (u *MatchPlayer) UpdateScore(score uint64) {
	u.Score.Score = score
}
