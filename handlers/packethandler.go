package handlers

import (
	"bancho/packets"
	"io"
	"sync"
	"time"

	"bancho/common"
	"bancho/events"
	"bytes"
	"container/list"
	"runtime/debug"
	"bancho/common/log"
	"bancho/inbound"
)

type UserDataInfo struct {
	ID         int32
	PlayerName string
	UTCOffset  byte
	Country    byte
	Colour     byte
	Longitude  float32
	Latitude   float32
	Rank       uint32
}

func Handle(input []byte, output io.Writer, token string) (string, error) {
	defer func() {
		c := recover()
		if c != nil {
			log.Error("ERROR!!!!!!!11!")
			log.Error(c)
			log.Error(string(debug.Stack()))
		}
	}()
	var sendBackToken bool
	var self *common.Session
	if token == "" {
		sendBackToken = true
		token, _, _ = events.Login(input)
		self = common.GetSession(token)
	} else if self = common.GetSession(token); self == nil || self.User.ID == 0 {
		sendBackToken = true
		token = common.GenerateGUID()
		self = &common.Session{
			LastRequest: time.Now(),
			Stream:      list.New(),
			Mutex:       &sync.Mutex{},
		}
		common.SessionsMutex.Lock()
		common.Sessions[token] = self
		common.SessionsMutex.Unlock()
		self.Push(
			packets.OrangeNotification("yo"),
			packets.UserID(-5),
		)
	} else {
		inputReader := bytes.NewReader(input)
		for {
			// Find a new packet from input
			pack, err := inbound.GetPacket(inputReader)
			if err != nil && err != io.EOF {
				log.Error(err)
			}
			if !pack.Initialised {
				break
			}
			ps := common.PackSess{
				P: pack,
				S: self,
			}
			HandleEvent(ps)
		}
	}

	self.Mutex.Lock()
	var e *list.Element
	for {
		e = self.Stream.Front()
		if e == nil {
			break
		}
		if actualE, can := e.Value.([]byte); can {
			output.Write(actualE)
		}
		self.Stream.Remove(e)
	}
	self.Mutex.Unlock()

	if sendBackToken {
		return token, nil
	}
	return "", nil
}
