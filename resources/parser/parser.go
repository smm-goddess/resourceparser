package parser

import (
	"bytes"
	"fmt"
	. "github.com/smm-goddess/resourceparser/resources/items"
)

func Parse(bs []byte) {
	buffer := bytes.NewBuffer(bs)
	tableHeader := ParseTableHeader(buffer)
	fmt.Println(tableHeader)
	globalStringPool := ParseStringPool(buffer)
	fmt.Println(globalStringPool.StringCount)
	fmt.Println(globalStringPool.Strings[3])
	tablePackage := ParsePackage(buffer)
	var buff bytes.Buffer
	for _, c := range tablePackage.Name {
		if c != 0x00 {
			buff.WriteByte(c)
		}
	}
	fmt.Println(buff.String())
	fmt.Println(tablePackage.KeyStringPool.Strings)
	fmt.Println(tablePackage.TypeStringPool.Strings)
}
