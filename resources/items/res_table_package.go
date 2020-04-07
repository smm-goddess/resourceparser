package items

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ResTablePackage struct {
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
	Name [128]uint16
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
Type String Pool
Key String Pool
Type Specification
Type Info
*/

func ParsePackage(src []byte, start uint32) ResTablePackage {
	var resTablePackage ResTablePackage
	_ = binary.Read(bytes.NewBuffer(src[start:]), binary.LittleEndian, &resTablePackage)
	fmt.Println(resTablePackage.TypeStrings)
	fmt.Println(resTablePackage.KeyStrings)
	return resTablePackage
}
