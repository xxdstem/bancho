package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"fmt"
)

func HandlePublicMessage(ps common.PackSess) {
	var(
		message string
		destination string
	)
	err := ps.P.Unmarshal(&message, &message, &destination)
	if err != nil{
		log.Error(err)
	}
	s := common.GetStream(fmt.Sprintf("chat/#%s", destination))
	if s != nil{
		s.Send(packets.SendMessage(&ps.S.User, destination, message))
	}
	log.Debug("User %d: Public message: %s, dest: %s", ps.S.User.ID, message, destination)
}


func HandlePrivateMessage(ps common.PackSess) {
	var(
		message string
		destination string
	)
	err := ps.P.Unmarshal(&message, &message, &destination)
	if err != nil{
		log.Error(err)
	}
	log.Debug("User %d: Private message: %s, dest: %s", ps.S.User.ID, message, destination)
}
