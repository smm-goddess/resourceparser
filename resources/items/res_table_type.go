package items

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func getResId(packageId, resTypeId, entryId uint32) uint32 {
	return ((packageId) << 24) | (((resTypeId) & 0xFF) << 16) | (entryId & 0xFFFF)
}

func ReadEntry(src []byte, packageId, resTypeId, offset, entryId uint32, globalString []string) {
	buffer := bytes.NewBuffer(src[offset:])
	getResId(packageId, resTypeId, entryId)
	//resId := getResId(packageId, resTypeId, entryId)
	//fmt.Printf("resId:%x ", resId)
	var entry ResTableEntry
	_ = binary.Read(buffer, binary.LittleEndian, &entry)
	if entry.Flags == 1 { // ResTableMapEntry
		var resTableMapEntry ResTableMapEntry
		_ = binary.Read(buffer, binary.LittleEndian, &resTableMapEntry)
		for j := uint32(0); j < resTableMapEntry.Count; j++ {
			var resTableMap ResTableMap
			_ = binary.Read(buffer, binary.LittleEndian, &resTableMap)
			if resTableMap.Value.DataType == TYPE_STRING {
				fmt.Println("string value map:", globalString[resTableMap.Value.Data])
			}
		}
	} else { // ResValue
		var resValue ResValue
		_ = binary.Read(buffer, binary.LittleEndian, &resValue)
		if resValue.DataType == TYPE_STRING {
			fmt.Println("string value:", globalString[resValue.Data])
		}
	}
}
