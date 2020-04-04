package items

import (
	"bytes"
	"encoding/binary"
)

type ResTablePackageHeader struct {
	/*
		chunk header
	*/
	ResChunkHeader
	/*
		package id, user resource package id = 0x7F, system resource package id: 0x01
	*/
	Id uint32
	/*
		package name
	*/
	Name [128]byte
	/*
		type string pool start place relative to header
	*/
	TypeStrings uint32
	/*
		last public string's index. equals to string pool's size at this time and not used during parse
	*/
	LastPublicType uint32
	/*
	 resource name string start place relative to header
	*/
	KeyStrings uint32
	/*
		last public key's index. equals to key pool's size at this point and not used during parse
	*/
	LastPublicKey uint32
}

/*
Package chunk data struct
String Pool
Type String Pool
Key String Pool
Type Specification
Type Info
*/

type ResTablePackage struct {
	ResTablePackageHeader
	/*
		struct same as string pool
	*/
	TypeStringPool ResStringPool
	/*
		struct same as string pool
	*/
	KeyStringPool     ResStringPool
	TypeSpecification uint32
	TypeInfo          uint32
}

func ParsePackage(buffer *bytes.Buffer) ResTablePackage {
	var tablePackageHeader ResTablePackageHeader
	_ = binary.Read(buffer, binary.LittleEndian, &tablePackageHeader)
	typeStringPool := ParseStringPool(buffer)
	KeyStringPool := ParseStringPool(buffer)
	return ResTablePackage{
		ResTablePackageHeader: tablePackageHeader,
		TypeStringPool:        typeStringPool,
		KeyStringPool:         KeyStringPool,
		TypeSpecification:     0,
		TypeInfo:              0,
	}
}
