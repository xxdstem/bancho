package common

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

var lastMatchID int

var Sessions map[string]*Session
var SessionsMutex *sync.RWMutex

var Matches map[int]*Match
var MatchesMutex *sync.RWMutex

var UidToSession map[int32]*Session
var UidToSessionMutex *sync.RWMutex

var UsernameToSession map[string]*Session
var UsernameToSessionMutex *sync.RWMutex

func Init() {
	SessionsMutex = &sync.RWMutex{}
	UidToSessionMutex = &sync.RWMutex{}
	UsernameToSessionMutex = &sync.RWMutex{}

	Sessions = make(map[string]*Session)
	UidToSession = make(map[int32]*Session)
	UsernameToSession = make(map[string]*Session)

	streams = make(map[string]*Stream)
	streamsMutex = &sync.RWMutex{}

	Matches = make(map[int]*Match)
	MatchesMutex = &sync.RWMutex{}
	NewStream("main")
	NewStream("lobby")

	// TODO: Initialize chat streams from DB
	NewStream("chat/#osu")
	NewStream("chat/#lobby")
}
