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
	User        *User
	LastRequest time.Time
}

type PackSess struct {
	P inbound.BasePacket
	S *Session
}
