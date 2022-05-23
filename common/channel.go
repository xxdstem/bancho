package common

import "bancho/packets"

type Channel struct {
	ID          uint32
	Name        string
	ClientName  string
	Description string
	//TODO: Add permission property's
	Temporary bool
	ReadOnly  bool // Moderated property in old Pep.py
	Stream    *Stream
}

func (c *Channel) SendMessage(sender *User, message string) {
	packet := packets.SendMessage(sender.Name, sender.ID, c.Name, message)
	packet.Ignored = append(packet.Ignored, sender.Token)
	c.Stream.Send(packet)
}
