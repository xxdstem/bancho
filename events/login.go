package events

import (
	"errors"
	"strconv"
	"strings"
	"github.com/xxdstem/bancho/packets"
	"github.com/xxdstem/bancho/common"
	
)

// LoginData is the data received by the osu! client upon a login request to bancho.

func Login(input []byte) (string, bool, error){
	sess, guid := common.NewSession(common.User{})
	loginData, err := Unmarshal(input)
	if err != nil{
		sess.Push(packets.UserID(-1))
	}
	err = common.DB.QueryRow("SELECT id, username FROM users WHERE username LIKE ?", loginData.Username).Scan(&sess.User.ID, &sess.User.Name)
	if err != nil{
		sess.Push(packets.UserID(-1))
	}
	sess.User.UpdateStats(0)
	sess.Push(
		packets.SilenceEnd(0),
		packets.UserID(sess.User.ID),
		packets.ChoProtocol(19),
		packets.UserPrivileges(),
		packets.FriendList([]int32{0}),
		packets.UserData(&sess.User),
		packets.UserDataFull(&sess.User),
		packets.OnlinePlayers(),
		packets.ChannelJoin(),
		packets.ChannelInfo(),
	)
	sess.Push(packets.ChannelListingComplete())
	
	common.UidToSessionMutex.Lock()
	common.UidToSession[int32(sess.User.ID)] = sess
	common.UidToSessionMutex.Unlock()
	s := common.GetStream("all")
	s.Subscribe(guid)
	go sendUserPresence(s, int32(sess.User.ID))
	return guid, false, nil
}

func sendUserPresence(s *common.Stream, uid int32) {
	count := 0
	for _, session := range common.CopySessions() {
		if session.User.ID == uid {
			count++
		}
	}
	if count < 2 {
		s.Send(packets.UserPresence(uid))
	}
}

// Unmarshal creates a new LoginData with the data passed.
func Unmarshal(input []byte) (l common.LoginData, e error) {
	lines := strings.Split(string(input), "\n")
	if len(lines) != 4 {
		e = errors.New("logindata: cannot unmarshal, got " + strconv.Itoa(len(lines)) + " lines as an input, want 4")
		return
	}
	l.Username = lines[0]
	l.Password = lines[1]
	l.HardwareData = strings.Split(lines[2], "|")
	l.HardwareHashes = strings.Split(lines[3], ":")
	return
}
