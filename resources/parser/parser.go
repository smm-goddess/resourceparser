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
	if resTablePackage.KeyStrings != 0 {
		packageKeyStrings := ParseStringPool(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size+resTablePackage.KeyStrings)
		offset += packageKeyStrings.Size
	}
	if resTablePackage.TypeStrings != 0 {
		packageTypeStrings := ParseStringPool(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size+resTablePackage.TypeStrings)
		offset += packageTypeStrings.Size
	}
	fmt.Printf("%x \n", offset)
}
