package events

import (
	"github.com/xxdstem/bancho/common"
	"fmt"
)

func PartMatch(ps common.PackSess){
	ps.S.User.LeaveMatch()
	UpdateMatch(ps)
	fmt.Println("leaving match!")
}