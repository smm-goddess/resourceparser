package parser

import (
	"bytes"
	"fmt"

	. "github.com/smm-goddess/resourceparser/resources/items"
)

func Parse(src []byte) {
	tableHeader := ParseTableHeader(src)
	globalStringPool := ParseStringPool(src, uint32(tableHeader.HeaderSize))
	fmt.Println("global string pool count:", globalStringPool.StringCount)
	resTablePackage := ParsePackage(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size)
	var buff bytes.Buffer
	for _, c16 := range resTablePackage.Name {
		if c16 != 0x00 {
			buff.WriteByte(uint8(c16))
		}
	}
	offset := uint32(tableHeader.HeaderSize+resTablePackage.HeaderSize) + globalStringPool.Size
	var packageKeyStrings, packageTypeStrings ResStringPool
	if resTablePackage.KeyStrings != 0 {
		packageKeyStrings = ParseStringPool(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size+resTablePackage.KeyStrings)
		offset += packageKeyStrings.Size
	}
	if resTablePackage.TypeStrings != 0 {
		packageTypeStrings = ParseStringPool(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size+resTablePackage.TypeStrings)
		offset += packageTypeStrings.Size
	}
	var resTableTypeSpec ResTableTypeSpec
	var resTableType ResTableType
	for offset < tableHeader.Size {
		if src[offset] == 1 && src[offset+1] == 2 { // ResTableType
			resTableType = ParseResTableType(src, offset)
			entriesOffset := ReadEntries(src, uint32(resTableType.ResChunkHeader.HeaderSize)+offset, resTableType.EntryCount)
			for entryId, diff := range entriesOffset {
				if diff != NO_ENTRY {
					ReadEntry(src, resTablePackage.Id, uint32(resTableType.Id), offset+diff, uint32(entryId), globalStringPool.Strings)
				}
			}
			offset += resTableType.ResChunkHeader.Size
		} else if src[offset] == 2 && src[offset+1] == 2 { // ResTableTypeSpec
			resTableTypeSpec = ParseResTableTypeSpec(src, offset)
			fmt.Println("int elements:", resTableTypeSpec.EntryCount)
			_ = ReadEntries(src, offset+uint32(resTableTypeSpec.ResChunkHeader.HeaderSize), resTableTypeSpec.EntryCount)
			offset += resTableTypeSpec.ResChunkHeader.Size
		} else {
			fmt.Println("else")
		}

		//fmt.Println(packageTypeStrings.Strings[resTableTypeSpec.Id-1])
		//fmt.Println(uint32(resTableTypeSpec.HeaderSize)+resTableTypeSpec.EntryCount*4, resTableTypeSpec.Size)

		//for _, entry := range entries {
		//	if entry != 0 {
		//		fmt.Println(entry)
		//	}
		//}
	}
}
