package parser

import (
	"bytes"
	"fmt"

	. "github.com/smm-goddess/resourceparser/resources/items"
)

func Parse(src []byte) {
	tableHeader := ParseTableHeader(src)
	globalStringPool := ParseStringPool(src, uint32(tableHeader.HeaderSize))
	fmt.Println(globalStringPool.StringCount)
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
			fmt.Println(resTableType)
			fmt.Printf("entry start %x \n", resTableType.EntriesStart)
			//fmt.Println("table type:", resTableType.EntryCount)
			//fmt.Println(packageTypeStrings.Strings[resTableType.Id-1])
			//fmt.Printf("%x \n", resTableType.ResTableConfig)
			fmt.Printf("%d \n", resTableType.HeaderSize)
			fmt.Printf("%d \n", resTableType.ResTableConfig.Size)
			//entriesOffset := ReadEntries(src, uint32(resTableType.HeaderSize)+offset, resTableType.EntryCount)
			//for _, offset := range entriesOffset {
			//	fmt.Println(offset)
			//}
			offset += resTableType.ResChunkHeader.Size
		} else if src[offset] == 2 && src[offset+1] == 2 { // ResTableTypeSpec
			resTableTypeSpec = ParseResTableTypeSpec(src, offset)
			//_ = ReadEntries(src, offset+uint32(resTableTypeSpec.HeaderSize), resTableTypeSpec.EntryCount)
			//fmt.Println("table spec:", resTableTypeSpec.EntryCount)
			fmt.Println(packageTypeStrings.Strings[resTableTypeSpec.Id-1])
			offset += resTableTypeSpec.Size
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
