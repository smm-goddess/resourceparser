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

func ParseTableHeader(buffer *bytes.Buffer) ResTableHeader {
	var resTableHeader ResTableHeader
	_ = binary.Read(buffer, binary.LittleEndian, &resTableHeader)
	return resTableHeader
}

func (this ResTableHeader) String() string {
	return fmt.Sprintf("%s, packageCount: %d", this.ResChunkHeader.String(), this.PackageCount)
}
