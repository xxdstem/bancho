package events

import (
	"github.com/xxdstem/bancho/common"
	"github.com/xxdstem/bancho/packets"
	"fmt"
)

func PartMatch(ps common.PackSess){
	matchID := ps.S.User.Match.ID
	dispose := ps.S.User.LeaveMatch()
	if dispose{
		s := common.GetStream("lobby")
		s.Send(packets.DisposeMatch(matchID))
		fmt.Println("pushing dispose match")
	}else{
		fmt.Println("just updatin")
		UpdateMatch(ps)
	}
	
	fmt.Println("leaving match!")
}