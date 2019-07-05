package packets

import (
	"github.com/xxdstem/bancho/common"
	"fmt"
)

func LoginFailed() common.FinalPacket{
	//return MakePacketOLD(5, 4, -1)
	return MakePacket(5, []Packet{{-1, SINT32}})
}

func ForceUpdate() common.FinalPacket{
	return MakePacket(5, []Packet{{-2, SINT32}})
}

func LoginError() common.FinalPacket{
	return MakePacket(5, []Packet{{-5, SINT32}})
}

func UserID(userID int32) common.FinalPacket{
	return MakePacket(5, []Packet{{userID, SINT32}})
}

func SilenceEnd(seconds uint32) common.FinalPacket{
	return MakePacket(92, []Packet{{seconds, UINT32}})
}

func ChoProtocol(version uint32) common.FinalPacket{
	return MakePacket(75, []Packet{{version, UINT32}})
}

func UserPrivileges() common.FinalPacket{
	return MakePacket(71, []Packet{{4, UINT32}})
}

func FriendList(friends []int32) common.FinalPacket {
	return MakePacket(72, []Packet{{friends, INT_LIST}})
}

func OnlinePlayers() common.FinalPacket {
	users := make([]int32, len(common.Sessions)+1)
	users[0] = 999
	i := 1
	for _, sess := range common.CopySessions() {
		if sess != nil && sess.User.ID != 0 {
			if i >= len(users) {
				users = append(users, sess.User.ID)
			} else {
				users[i] = sess.User.ID
			}
			i++
		}
	}
	fmt.Println(users[:i])
	return MakePacket(96, []Packet{{users[:i], INT_LIST}})
}

func ChannelJoin() common.FinalPacket{
	return MakePacket(64, []Packet{
		{"#osu", STRING},
	})
}

func ChannelInfo() common.FinalPacket{
	return MakePacket(65, []Packet{
		{"#osu", STRING},
		{"Main channel", STRING},
		{1, UINT16},
	})
}

func ChannelListingComplete() common.FinalPacket {
	return MakePacket(89, []Packet{{0, UINT32}})
}

func UserData(user *common.User) common.FinalPacket{
	packetData := []Packet{
		{user.ID, SINT32},
		{user.Name, STRING},
		{27, BYTE},
		{56, BYTE},
		{0, BYTE},
		{0.0, FLOAT},
		{0.0, FLOAT}, 
		{0, UINT32}, //rank?
	}
	return MakePacket(83, packetData)
}

func BotData() common.FinalPacket{
	packetData := []Packet{
		{999, SINT32},
		{"FokaBot", STRING},
		{27, BYTE},
		{0, BYTE},
		{0, BYTE},
		{0.0, FLOAT},
		{0.0, FLOAT}, 
		{0, UINT32},
	}
	return MakePacket(83, packetData)
}

func UserDataFull(user *common.User) common.FinalPacket{
	packetData := []Packet{
		{user.ID, SINT32},
		{user.Status.Status, BYTE}, //a id
		{user.Status.Text, STRING}, //a text
		{user.Status.MD5, STRING}, //a md5
		{user.Status.Mods, SINT32}, //mods
		{user.Stats.Mode, BYTE}, //gm
		{user.Status.BeatmapID, SINT32}, //bid
		{user.Stats.RankedScore, UINT64}, // rankedscore
		{user.Stats.Accuracy, FLOAT}, //accuracy
		{user.Stats.PlayCount, UINT32}, // playcount
		{user.Stats.TotalScore, UINT64}, // totalScore
		{user.Stats.Rank, UINT32}, // gameRank
		{user.Stats.PP, UINT16}, // pp

	}
	return MakePacket(11, packetData)
}

func UserPresence(userID int32) common.FinalPacket {
	return MakePacket(85, []Packet{{userID, SINT32}})
}	

func OrangeNotification(message string) common.FinalPacket {
	return MakePacket(24, []Packet{{message, STRING}})
}

func MatchDataFull(m *common.Match, packetID uint16, censored bool) common.FinalPacket{
	var password string
	if censored && m.Password != ""{
		password = "redacted"
	}else{
		password = m.Password
	}
	pack := []Packet{
		{uint16(m.ID), UINT16},
		{byte(0), BYTE},
		{byte(0), BYTE},
		{uint32(0), UINT32},
		{m.Name, STRING},
		{password, STRING},
		{m.Beatmap.Name, STRING},
		{m.Beatmap.ID, UINT32},
		{m.Beatmap.MD5, STRING},
	}
	for _, slot := range m.Players{
		pack = append(pack, Packet{slot.Status, BYTE})
	}
	for _, slot := range m.Players{
		pack = append(pack, Packet{slot.Team, BYTE})
	}

	pack = append(pack,
		Packet{uint32(m.HostID), UINT32},

		Packet{m.HostID, SINT32},
		Packet{0, BYTE},
		Packet{0, BYTE},
		Packet{0, BYTE},
		Packet{0, UINT32},
	)
	return MakePacket(packetID, pack)
}

func DisposeMatch(matchID uint32) common.FinalPacket{
	return MakePacket(BanchoMatchDisband, []Packet{{matchID, UINT32}})
}