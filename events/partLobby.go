package events

import (
	"github.com/xxdstem/bancho/common"
	"fmt"
)

func PartLobby(ps common.PackSess){
	s := common.GetStream("lobby")
	s.Unsubscribe(ps.S.User.Token)
	fmt.Println("part lobby :(")
}