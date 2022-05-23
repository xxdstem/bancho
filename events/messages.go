package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
	"bancho/handlers/bot"
	"bancho/packets"
)

func HandlePublicMessage(ps common.PackSess) {
	var (
		message     string
		destination string
	)
	err := ps.P.Unmarshal(&message, &message, &destination)
	if err != nil {
		log.Error(err)
		return
	}

	ch := chat.GetChannel(destination)

	if ch == nil {
		log.Error("Empty channel %s?", destination)
		return
	}

	ch.SendMessage(ps.S.User, message)

	log.Debug("User %d: Public message: %s, dest: %s", ps.S.User.ID, message, destination)
}

func HandlePrivateMessage(ps common.PackSess) {
	var (
		message     string
		destination string
		packet      packets.FinalPacket
	)
	err := ps.P.Unmarshal(&message, &message, &destination)
	if err != nil {
		log.Error(err)
	}
	sess := common.GetSessionByUsername(common.SafeUsername(destination))
	if sess.User.ID == 999 {
		msg := bot.HandleMessage(message)
		if msg != "" {
			packet = packets.SendMessage("GoBot", 999, ps.S.User.Name, msg)
			ps.S.Push(packet)
		}
		return
	}
	packet = packets.SendMessage(ps.S.User.Name, ps.S.User.ID, destination, message)
	sess.Push(packet)
}
