package common

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

var lastMatchID int

var Sessions map[string]*Session = make(map[string]*Session)
var SessionsMutex *sync.RWMutex = &sync.RWMutex{}

var Matches map[int]*Match = make(map[int]*Match)
var MatchesMutex *sync.RWMutex = &sync.RWMutex{}

var UidToSession map[int32]*Session = make(map[int32]*Session)
var UidToSessionMutex *sync.RWMutex = &sync.RWMutex{}

var UsernameToSession map[string]*Session = make(map[string]*Session)
var UsernameToSessionMutex *sync.RWMutex = &sync.RWMutex{}

func Init() {
	NewStream("main")
	NewStream("lobby")
	botSess, _ := NewSession(User{
		ID:       999,
		Name:     "GoBot",
		SafeName: "gobot",
		mutex:    &sync.RWMutex{},
	})
	UsernameToSession[botSess.User.SafeName] = botSess
	UidToSession[int32(botSess.User.ID)] = botSess
	// TODO: Initialize chat streams from DB

}
