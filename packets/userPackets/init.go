package userPackets

import (
	"bancho/common"
	"bancho/packets"
)

func UserData(user *common.User) packets.FinalPacket {
	packetData := []packets.Packet{
		{user.ID, packets.SINT32},
		{user.Name, packets.STRING},
		{27, packets.BYTE},
		{56, packets.BYTE},
		{0, packets.BYTE},
		{0.0, packets.FLOAT},
		{0.0, packets.FLOAT},
		{0, packets.UINT32}, //rank?
	}
	return packets.MakePacket(83, packetData)
}

func OnlinePlayers() packets.FinalPacket {
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
	return packets.MakePacket(96, []packets.Packet{{users[:i], packets.INT_LIST}})
}

func UserDataFull(user *common.User) packets.FinalPacket {
	packetData := []packets.Packet{
		{user.ID, packets.SINT32},
		{user.Status.Status, packets.BYTE},       //a id
		{user.Status.Text, packets.STRING},       //a text
		{user.Status.MD5, packets.STRING},        //a md5
		{user.Status.Mods, packets.SINT32},       //mods
		{user.Stats.Mode, packets.BYTE},          //gm
		{user.Status.BeatmapID, packets.SINT32},  //bid
		{user.Stats.RankedScore, packets.UINT64}, // rankedscore
		{user.Stats.Accuracy, packets.FLOAT},     //accuracy
		{user.Stats.PlayCount, packets.UINT32},   // playcount
		{user.Stats.TotalScore, packets.UINT64},  // totalScore
		{user.Stats.Rank, packets.UINT32},        // gameRank
		{user.Stats.PP, packets.UINT16},          // pp

	}
	return packets.MakePacket(11, packetData)
}

func MatchDataFull(m *common.Match, packetID uint16, censored bool) packets.FinalPacket {
	var password string
	if censored && m.Password != "" {
		password = "redacted"
	} else {
		password = m.Password
	}
	pack := []packets.Packet{
		{uint16(m.ID), packets.UINT16},
		{byte(0), packets.BYTE},
		{byte(0), packets.BYTE},
		{uint32(0), packets.UINT32},
		{m.Name, packets.STRING},
		{password, packets.STRING},
		{m.Beatmap.Name, packets.STRING},
		{m.Beatmap.ID, packets.UINT32},
		{m.Beatmap.MD5, packets.STRING},
	}
	for _, slot := range m.Players {
		pack = append(pack, packets.Packet{slot.Status, packets.BYTE})
	}
	for _, slot := range m.Players {
		pack = append(pack, packets.Packet{slot.Team, packets.BYTE})
	}

	pack = append(pack,
		packets.Packet{uint32(m.HostID), packets.UINT32},

		packets.Packet{m.HostID, packets.SINT32},
		packets.Packet{0, packets.BYTE},
		packets.Packet{0, packets.BYTE},
		packets.Packet{0, packets.BYTE},
		packets.Packet{0, packets.UINT32},
	)
	return packets.MakePacket(packetID, pack)
}
