package common

import (
	"container/list"
	"sync"
	"time"

	"bancho/inbound"
)

type LoginData struct {
	Username       string
	Password       string
	HardwareData   []string
	HardwareHashes []string
}

// Session is an alive connection of a logged in user.
type Session struct {
	Stream      *list.List
	Mutex       *sync.Mutex
	User        User
	LastRequest time.Time
}

type FinalPacket struct {
	Content []byte
	// Ignored is a series of users of which this packet should NEVER arrive.
	Ignored []string
}

type PackSess struct {
	P inbound.BasePacket
	S *Session
}
