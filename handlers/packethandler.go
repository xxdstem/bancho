package handlers

import (
	"github.com/xxdstem/bancho/packets"
	"io"
	"fmt"
	"time"
	"sync"

	"runtime/debug"
	"container/list"
	"bytes"
	"git.zxq.co/ripple/nuclearbancho/inbound"
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/events"
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
func Handle(input []byte, output io.Writer, token string) (string, error){
	defer func() {
		c := recover()
		if c != nil {
			fmt.Println("ERROR!!!!!!!11!")
			fmt.Println(c)
			fmt.Println(string(debug.Stack()))
		}
	}()
	var sendBackToken bool
	var self *common.Session
	if token == "" {
		sendBackToken = true
		token, _, _ = events.Login(input)
		self = common.GetSession(token)
	}else if self = common.GetSession(token); self == nil || self.User.ID == 0 {
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
	}else{
		inputReader := bytes.NewReader(input)
		for {
			// Find a new packet from input
			pack, err := inbound.GetPacket(inputReader)
			if err != nil && err != io.EOF {
				fmt.Println(err)
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


