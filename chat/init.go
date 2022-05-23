package chat

import (
	"bancho/common"
	"fmt"
	"strings"
	"sync"
)

var channels map[string]*common.Channel = make(map[string]*common.Channel)
var channelsMutex *sync.RWMutex = &sync.RWMutex{}

// TODO: Make IRC Server messaging here

func NewChannel(name string) {
	channelsMutex.Lock()
	stream := common.NewStream(fmt.Sprintf("chat/#%s", name))
	channels[name] = &common.Channel{
		ID:         1,
		Name:       name,
		ClientName: ConvertToClientName(name),
		Stream:     stream,
	}
	channelsMutex.Unlock()
}

func GetChannel(name string) *common.Channel {
	if ch, ok := channels[name]; ok {
		return ch
	}
	return nil
}

func ConvertToClientName(name string) string {
	if strings.HasPrefix(name, "#spect_") {
		return "#spectator"
	}
	if strings.HasPrefix(name, "#multi_") {
		return "#multiplayer"
	}

	return name
}

func init() {

	NewChannel("#osu")
	NewChannel("#lobby")

}
