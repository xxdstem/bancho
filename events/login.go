package events

import (
	"bancho/common"
	"bancho/common/log"
	"bancho/packets"
	"errors"
	"strconv"
	"strings"
)

// LoginData is the data received by the osu! client upon a login request to bancho.

func Login(input []byte) (string, bool, error) {
	sess, guid := common.NewSession(common.User{})
	loginData, err := Unmarshal(input)
	if err != nil {
		sess.Push(packets.UserID(-1))
	}
	err = common.DB.QueryRow("SELECT id, username FROM users WHERE username LIKE ?", loginData.Username).Scan(&sess.User.ID, &sess.User.Name)
	if err != nil {
		sess.Push(packets.UserID(-1))
	}
	sess.User.UpdateStats(0)
	sess.Push(
		packets.SilenceEnd(0),
		packets.UserID(sess.User.ID),
		packets.ChoProtocol(19),
		packets.UserPrivileges(),
		packets.FriendList([]int32{0}),
		packets.UserData(sess.User),
		packets.UserDataFull(sess.User),
		packets.OnlinePlayers(),
		packets.ChannelJoin(),
		packets.ChannelInfo(),
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
	chat := common.GetStream("chat/#osu")
	if chat == nil {
		log.Error("for some reason #osu chat is unknown!")
	}
	chat.Subscribe(guid)

	go main.Send(packets.UserPresence(int32(sess.User.ID)))
	go main.Send(packets.UserData(sess.User))
	go sendPlayersStats(sess)
	return guid, false, nil
}

func sendPlayersStats(s *common.Session) {
	for _, session := range common.CopySessions() {
		if session.User.ID != s.User.ID {
			s.Push(packets.UserData(session.User))
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
