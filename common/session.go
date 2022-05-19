package common

import (
	"container/list"
	"sync"
	"time"
)

// Push appends an element to the current session.
func (s Session) Push(val ...FinalPacket) {
	//dumper := banchoreader.New()
	//dumper.Colored = true
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	for _, v := range val {
		var c bool
		for _, ignored := range v.Ignored {
			if s.User.Token == ignored {
				c = true
				break
			}
		}
		if c {
			continue
		}
		s.Stream.PushBack(v.Content)
		//dumper.Dump(os.Stdout, v.Content)
	}
}

// NewSession generates a new session.
func NewSession(u User) (*Session, string) {
	var tok string
	for {
		tok = GenerateGUID()
		// Make sure token does not already exist
		if _, ok := Sessions[tok]; !ok {
			break
		}
	}
	u.Token = tok
	sess := &Session{
		Stream:      list.New(),
		Mutex:       &sync.Mutex{},
		User:        u,
		LastRequest: time.Now(),
	}
	SessionsMutex.Lock()
	defer SessionsMutex.Unlock()
	Sessions[tok] = sess
	return sess, tok
}

func DeleteSession(s *Session) error {

	if ses, ok := Sessions[s.User.Token]; ok {
		delete(Sessions, ses.User.Token)
		delete(UsernameToSession, s.User.SafeName)
		delete(UidToSession, s.User.ID)
	}
	return nil
}

// GetSession retrieves a session from the available ones.
func GetSession(sessName string) *Session {
	SessionsMutex.RLock()
	defer SessionsMutex.RUnlock()
	return Sessions[sessName]
}

// CopySessions can be used to get an independent copy of Sessions, without need to use the sessionMutex to modify it.
func CopySessions() map[string]*Session {
	SessionsMutex.RLock()
	defer SessionsMutex.RUnlock()
	ret := make(map[string]*Session, len(Sessions))
	for k, v := range Sessions {
		ret[k] = v
	}
	return ret
}

// GetSessionByID returns a session retrieving it using his ID.
func GetSessionByID(id int32) *Session {
	UidToSessionMutex.RLock()
	defer UidToSessionMutex.RUnlock()
	v, _ := UidToSession[id]
	return v
}

// GetSessionByUsername returns a session retrieving it using his username
func GetSessionByUsername(username string) *Session {
	UsernameToSessionMutex.RLock()
	defer UsernameToSessionMutex.RUnlock()
	v, _ := UsernameToSession[username]
	return v
}
