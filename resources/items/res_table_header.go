package items

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ResTableHeader struct {
	/*
		chunk header
	*/
	ResChunkHeader
	/*
		package count in this package
	*/
	PackageCount uint32
}

func ParseTableHeader(src []byte) ResTableHeader {
	var resTableHeader ResTableHeader
	_ = binary.Read(bytes.NewBuffer(src), binary.LittleEndian, &resTableHeader)
	return resTableHeader
}

func (this ResTableHeader) GetSize() uint32 {
	return this.ResChunkHeader.GetSize() + 4
}

func (this ResTableHeader) String() string {
	return fmt.Sprintf("%s, packageCount: %d", this.ResChunkHeader.String(), this.PackageCount)
}
