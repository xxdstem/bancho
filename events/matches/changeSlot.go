package events

import (
	"github.com/xxdstem/bancho/common"
	_"github.com/xxdstem/bancho/packets"
	"fmt"
)

func ChangeSlot(ps common.PackSess){
	match := ps.S.User.Match
	var slotID uint32
	ps.P.Unmarshal(&slotID)
	fmt.Println(ps.S.User.ID, "changed slot to ", slotID)
	match.Mutex.Lock()
	defer match.Mutex.Unlock()
	for id, slot := range match.Players{
		if slot.User != nil && slot.User.ID == ps.S.User.ID{
			match.Players[slotID] = match.Players[id]
			match.Players[id].Status = 1
			match.Players[id].User = nil
			UpdateMatch(ps)
			return
		}
	}
	
}