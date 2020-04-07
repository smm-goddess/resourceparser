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
	fmt.Printf("%x\n", uint32(tableHeader.HeaderSize)+globalStringPool.Size)
	tablePackage := ParsePackage(src, uint32(tableHeader.HeaderSize)+globalStringPool.Size)
	var buff bytes.Buffer
	for _, c := range tablePackage.Name {
		if c != 0x00 {
			buff.WriteByte(c)
		}
	}
	fmt.Println(buff.String())
}
