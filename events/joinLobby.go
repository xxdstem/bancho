package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"fmt"
)

func JoinLobby(ps common.PackSess){
	s := common.GetStream("lobby")
	s.Subscribe(ps.S.User.Token)
	common.MatchesMutex.Lock()
	defer common.MatchesMutex.Unlock()
	for _, v := range common.Matches{
		if v != nil{
			ps.S.Push(packets.MatchDataFull(v, packets.BanchoMatchNew, true))
		}
	}
	fmt.Println("joined lobby :)")
}