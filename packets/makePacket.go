package packets

import (
	//"bancho/common"

	"bytes"
	"encoding/binary"

	"github.com/bnch/uleb128"
)

func MakePacket(t uint16, packets []Packet) FinalPacket {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, t)
	binary.Write(b, binary.LittleEndian, byte(0))

	dataBuffer := new(bytes.Buffer)
	for _, v := range packets {
		if v.Type == STRING {
			conv, _ := v.Data.(string)
			binary.Write(dataBuffer, binary.LittleEndian, BanchoString(conv))
		} else if v.Type == SINT32 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, int32(conv))
			}

		} else if v.Type == UINT32 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, uint32(conv))
			}
		} else if v.Type == SINT16 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, int16(conv))
			}
		} else if v.Type == UINT16 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, uint16(conv))
			}
		} else if v.Type == SINT64 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, int64(conv))
			}
		} else if v.Type == UINT64 {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, uint64(conv))
			}
		} else if v.Type == FLOAT {
			conv, ok := v.Data.(float64)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, float32(conv))
			}
		} else if v.Type == BYTE {
			conv, ok := v.Data.(int)
			if !ok {
				binary.Write(dataBuffer, binary.LittleEndian, v.Data)
			} else {
				binary.Write(dataBuffer, binary.LittleEndian, byte(conv))
			}
		} else if v.Type == INT_LIST {
			conv, _ := v.Data.([]int32)
			binary.Write(dataBuffer, binary.LittleEndian, uint16(len(conv)))
			binary.Write(dataBuffer, binary.LittleEndian, conv)
		} else {
			binary.Write(dataBuffer, binary.LittleEndian, v.Data)
		}
	}
	endB := dataBuffer.Bytes()
	binary.Write(b, binary.LittleEndian, uint32(len(endB)))
	binary.Write(b, binary.LittleEndian, endB)
	return FinalPacket{
		Content: b.Bytes(),
	}
}

func BanchoString(s string) []byte {
	if s == "" {
		return []byte{0}
	}
	// 11, aka 0x0b, notifies the client that what's following is a string.
	r := []byte{11}
	r = append(r, uleb128.Marshal(len(s))...)
	r = append(r, []byte(s)...)
	return r
}
