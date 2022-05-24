package events

import (
	"bancho/chat"
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"bancho/packets/userPackets"
	"errors"
	"strconv"
	"strings"
)

// LoginData is the data received by the osu! client upon a login request to bancho.

func Login(input []byte) (string, bool, error) {
	u := common.User{
		Channels: make(map[string]*common.Channel),
	}
	var password string
	sess, guid := common.NewSession(u)
	loginData, err := Unmarshal(input)
	if err != nil {
		sess.Push(packets.UserID(-1))
		return guid, true, nil
	}
	err = common.DB.QueryRow("SELECT id, username, password_md5 FROM users WHERE username LIKE ?", loginData.Username).Scan(&sess.User.ID, &sess.User.Name, &password)
	if err != nil {
		sess.Push(packets.UserID(-1))
		return guid, true, nil
	}
	if !common.IsSamePass(loginData.Password, password) {
		sess.Push(packets.UserID(-1))
		return guid, true, nil
	}

	osuChannel := chat.GetChannel("#osu")
	sess.User.UpdateStats(0)
	sess.Push(
		packets.SilenceEnd(0),
		packets.UserID(sess.User.ID),
		packets.ChoProtocol(19),
		packets.UserPrivileges(),
		packets.FriendList([]int32{0}),
		userPackets.UserData(sess.User),
		userPackets.UserDataFull(sess.User),
		userPackets.OnlinePlayers(),
		userPackets.ChannelJoin(osuChannel),
		userPackets.ChannelInfo(osuChannel),
	)
	sess.Push(packets.ChannelListingComplete())

	common.UidToSessionMutex.Lock()
	common.UsernameToSessionMutex.Lock()

	common.UsernameToSession[sess.User.SafeName] = sess
	common.UidToSession[int32(sess.User.ID)] = sess

	common.UidToSessionMutex.Unlock()
	common.UsernameToSessionMutex.Unlock()

	main := common.GetStream("main")
	if main == nil {
		log.Error("niggers", main)
	}
	main.Subscribe(guid)

	sess.User.JoinChannel(osuChannel)

	go main.Send(packets.UserPresence(int32(sess.User.ID)))
	go main.Send(userPackets.UserData(sess.User))
	go sendPlayersStats(sess)
	return guid, false, nil
}

func sendPlayersStats(s *common.Session) {
	for _, session := range common.CopySessions() {
		if session.User.ID != s.User.ID {
			s.Push(userPackets.UserData(session.User))
		}
	}
}

// Unmarshal creates a new LoginData with the data passed.
func Unmarshal(input []byte) (l common.LoginData, e error) {
	lines := strings.Split(string(input), "\n")
	if len(lines) != 4 {
		e = errors.New("logindata: cannot unmarshal, got " + strconv.Itoa(len(lines)) + " lines as an input, want 4")
		return
	}
	l.Username = strings.TrimSpace(lines[0])
	l.Password = lines[1]
	l.HardwareData = strings.Split(lines[2], "|")
	l.HardwareHashes = strings.Split(lines[3], ":")
	return
}
