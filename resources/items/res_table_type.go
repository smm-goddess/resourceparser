package items

import (
	"bytes"
	"encoding/binary"
)

const (
	NO_ENTRY = 0xFFFFFFFF
)

/*
 type = 0x0202
*/
type ResTableType struct {
	ResChunkHeader
	/*
		type id
	*/
	Id   uint8
	Res0 uint8
	Res  uint16
	/*
		resource count
	*/
	EntryCount uint32
	/*
		resource data offset relative to header
	*/
	EntriesStart uint32
	/*
		ResTableConfig, locale, language, resolution etc.
	*/
	ResTableConfig
}

func ParseResTableType(src []byte, offset uint32) ResTableType {
	var resTableType ResTableType
	_ = binary.Read(bytes.NewBuffer(src[offset:]), binary.LittleEndian, &resTableType)
	return resTableType
}
