package events

import (
	"github.com/xxdstem/bancho/common"
	"fmt"
)

func JoinLobby(ps common.PackSess){
	s := common.GetStream("lobby")
	s.Subscribe(ps.S.User.Token)
	fmt.Println("joined lobby :)")
}