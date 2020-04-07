package items

import (
	"bytes"
	"encoding/binary"
)

const (
	SPEC_PUBLIC = uint32(0x40000000)
)

const (
	ACONFIGURATION_MCC                  = 0x0001
	ACONFIGURATION_MNC                  = 0x0002
	ACONFIGURATION_LOCALE               = 0x0004
	ACONFIGURATION_TOUCHSCREEN          = 0x0008
	ACONFIGURATION_KEYBOARD             = 0x0010
	ACONFIGURATION_KEYBOARD_HIDDEN      = 0x0020
	ACONFIGURATION_NAVIGATION           = 0x0040
	ACONFIGURATION_ORIENTATION          = 0x0080
	ACONFIGURATION_DENSITY              = 0x0100
	ACONFIGURATION_SCREEN_SIZE          = 0x0200
	ACONFIGURATION_VERSION              = 0x0400
	ACONFIGURATION_SCREEN_LAYOUT        = 0x0800
	ACONFIGURATION_UI_MODE              = 0x1000
	ACONFIGURATION_SMALLEST_SCREEN_SIZE = 0x2000
	ACONFIGURATION_LAYOUTDIR            = 0x4000
	ACONFIGURATION_SCREEN_ROUND         = 0x8000
	ONFIG_MCC                           = ACONFIGURATION_MCC
	CONFIG_MNC                          = ACONFIGURATION_MNC
	CONFIG_LOCALE                       = ACONFIGURATION_LOCALE
	CONFIG_TOUCHSCREEN                  = ACONFIGURATION_TOUCHSCREEN
	CONFIG_KEYBOARD                     = ACONFIGURATION_KEYBOARD
	CONFIG_KEYBOARD_HIDDEN              = ACONFIGURATION_KEYBOARD_HIDDEN
	CONFIG_NAVIGATION                   = ACONFIGURATION_NAVIGATION
	CONFIG_ORIENTATION                  = ACONFIGURATION_ORIENTATION
	CONFIG_DENSITY                      = ACONFIGURATION_DENSITY
	CONFIG_SCREEN_SIZE                  = ACONFIGURATION_SCREEN_SIZE
	CONFIG_SMALLEST_SCREEN_SIZE         = ACONFIGURATION_SMALLEST_SCREEN_SIZE
	CONFIG_VERSION                      = ACONFIGURATION_VERSION
	CONFIG_SCREEN_LAYOUT                = ACONFIGURATION_SCREEN_LAYOUT
	CONFIG_UI_MODE                      = ACONFIGURATION_UI_MODE
	CONFIG_LAYOUTDIR                    = ACONFIGURATION_LAYOUTDIR
	CONFIG_SCREEN_ROUND                 = ACONFIGURATION_SCREEN_ROUND
)

/*
	follow ResTableTypeSpec is a uint32 array, size equals entry count, each item is the diff of the item
*/
/*
 type = 0x0202
*/
type ResTableTypeSpec struct {
	ResChunkHeader
	/*
		resource-id:including: animator, anim, color, drawable, layout, raw, string, xml. each type has an id
	*/
	Id uint8
	/*
		not used = 0
	*/
	Res0 uint8
	/*
		not used = 0
	*/
	Res uint16
	/*
		this type of resource count
	*/
	EntryCount uint32
}

func ParseResTableTypeSpec(src []byte, offset uint32) ResTableTypeSpec {
	var resTableTypeSpec ResTableTypeSpec
	_ = binary.Read(bytes.NewBuffer(src[offset:]), binary.LittleEndian, &resTableTypeSpec)
	return resTableTypeSpec
}

func ReadEntries(src []byte, offset, count uint32) []uint32 {
	entries := make([]uint32, count, count)
	_ = binary.Read(bytes.NewBuffer(src[offset:]), binary.LittleEndian, &entries)
	return entries
}
