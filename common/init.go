package common

import (
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

var Sessions map[string]*Session
var SessionsMutex *sync.RWMutex


var Matches map[string]*Match
var MatchesMutex *sync.RWMutex

var UidToSession map[int32]*Session
var UidToSessionMutex *sync.RWMutex

func Init(){
	SessionsMutex = &sync.RWMutex{}
	UidToSessionMutex = &sync.RWMutex{}
	
	Sessions = make(map[string]*Session)
	UidToSession = make(map[int32]*Session)

	streams = make(map[string]*Stream)
	streamsMutex = &sync.RWMutex{}

	Matches = make(map[string]*Match)
	MatchesMutex = &sync.RWMutex{}
	NewStream("main")
	NewStream("lobby")
	
}