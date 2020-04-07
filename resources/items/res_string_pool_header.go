package items

import (
	"bytes"
	"encoding/binary"
)

const (
	SORTED_FLAG = 1
	UTF8_FLAG   = 1 << 8
)

type ResStringPoolHeader struct {
	/*
		standard chunk header
	*/
	ResChunkHeader
	/*
		string count
	*/
	StringCount uint32
	/*
		style count
	*/
	StyleCount uint32
	/*
		string flag, include 0x000(UTF-16),0x001(Sorted string),0x100(UTF-8)
	*/
	Flags uint32
	/*
		string content start position relative to header start
	*/
	StringStart uint32
	/*
		style content start position relative to header start
	*/
	StyleStart uint32
}

/*
string pool struct:
String Pool
ResStringPool Header
String Offset Array
Style Offset Array
Strings
Styles
*/

type ResStringPool struct {
	ResStringPoolHeader
	/*
		size equals string count
	*/
	StringOffsetArray []uint32
	/*
		size equals style count
	*/
	StyleOffsetArray []uint32
	Strings          []string
	Styles           []Style
}

type Style struct {
	ResStringPoolSpans []ResStringPoolSpan
	ResStringPoolRef
}

func ParseStringPool(src []byte, start uint32) ResStringPool {
	var stringPoolHeader ResStringPoolHeader
	buffer := bytes.NewBuffer(src[start:])
	_ = binary.Read(buffer, binary.LittleEndian, &stringPoolHeader)
	stringOffsetArray := make([]uint32, stringPoolHeader.StringCount, stringPoolHeader.StringCount)
	styleOffsetArray := make([]uint32, stringPoolHeader.StyleCount, stringPoolHeader.StyleCount)
	_ = binary.Read(buffer, binary.LittleEndian, &stringOffsetArray)
	_ = binary.Read(buffer, binary.LittleEndian, &styleOffsetArray)
	var stringBuff bytes.Buffer
	var stringLength uint16
	var b byte
	bs := make([]byte, 2, 2)
	strings := make([]string, stringPoolHeader.StringCount, stringPoolHeader.StringCount)
	buffer = bytes.NewBuffer(src[start+stringPoolHeader.StringStart:])
	for i := uint32(0); i < stringPoolHeader.StringCount; i++ {
		_ = binary.Read(buffer, binary.LittleEndian, &stringLength)
		if stringPoolHeader.Flags == UTF8_FLAG { // ends with 0x00
			b, _ = buffer.ReadByte()
			for b != 0x00 {
				stringBuff.WriteByte(b)
				b, _ = buffer.ReadByte()
			}
		} else { // ends with 0x0000
			_, _ = buffer.Read(bs)
			for bs[0] != 0x00 && bs[1] != 0x00 {
				stringBuff.Write(bs)
				_, _ = buffer.Read(bs)
			}
		}
		strings[i] = stringBuff.String()
		stringBuff.Reset()
	}
	// TODO parse style table;
	return ResStringPool{
		ResStringPoolHeader: stringPoolHeader,
		StringOffsetArray:   stringOffsetArray,
		StyleOffsetArray:    styleOffsetArray,
		Strings:             strings,
		Styles:              nil,
	}
}
