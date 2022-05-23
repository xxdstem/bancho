package packets

//"bancho/common"

func LoginFailed() FinalPacket {
	//return MakePacketOLD(5, 4, -1)
	return MakePacket(5, []Packet{{-1, SINT32}})
}

func ForceUpdate() FinalPacket {
	return MakePacket(5, []Packet{{-2, SINT32}})
}

func LoginError() FinalPacket {
	return MakePacket(5, []Packet{{-5, SINT32}})
}

func UserID(userID int32) FinalPacket {
	return MakePacket(5, []Packet{{userID, SINT32}})
}

func SilenceEnd(seconds uint32) FinalPacket {
	return MakePacket(92, []Packet{{seconds, UINT32}})
}

func ChoProtocol(version uint32) FinalPacket {
	return MakePacket(75, []Packet{{version, UINT32}})
}

func UserPrivileges() FinalPacket {
	return MakePacket(71, []Packet{{4, UINT32}})
}

func FriendList(friends []int32) FinalPacket {
	return MakePacket(72, []Packet{{friends, INT_LIST}})
}

//FULL REWORK

func ChannelJoin() FinalPacket {
	return MakePacket(64, []Packet{
		{"#osu", STRING},
	})
}

func ChannelInfo() FinalPacket {
	return MakePacket(65, []Packet{
		{"#osu", STRING},
		{"Main channel", STRING},
		{1, UINT16},
	})
}

func ChannelListingComplete() FinalPacket {
	return MakePacket(89, []Packet{{0, UINT32}})
}

func SendMessage(sender string, senderID int32, destination string, message string) FinalPacket {
	return MakePacket(BanchoSendMessage, []Packet{{sender, STRING}, {message, STRING}, {destination, STRING}, {senderID, SINT32}})
}

func BotData() FinalPacket {
	packetData := []Packet{
		{999, SINT32},
		{"GoBot", STRING},
		{27, BYTE},
		{0, BYTE},
		{0, BYTE},
		{0.0, FLOAT},
		{0.0, FLOAT},
		{0, UINT32},
	}
	return MakePacket(83, packetData)
}

func UserPresence(userID int32) FinalPacket {
	return MakePacket(85, []Packet{{userID, SINT32}})
}

func OrangeNotification(message string) FinalPacket {
	return MakePacket(24, []Packet{{message, STRING}})
}

func DisposeMatch(matchID uint32) FinalPacket {
	return MakePacket(BanchoMatchDisband, []Packet{{matchID, UINT32}})
}

func LogOut(userID int32) FinalPacket {
	return MakePacket(BanchoHandleUserQuit, []Packet{{userID, SINT32}, {0, BYTE}})
}
