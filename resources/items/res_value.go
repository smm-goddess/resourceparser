package items

const (
	TYPE_NULL            = 0x00
	TYPE_REFERENCE       = 0x01
	TYPE_ATTRIBUTE       = 0x02
	TYPE_STRING          = 0x03
	TYPE_FLOAT           = 0x04
	TYPE_DIMENSION       = 0x05
	TYPE_FRACTION        = 0x06
	TYPE_FIRST_INT       = 0x10
	TYPE_INT_DEC         = 0x10
	TYPE_INT_HEX         = 0x11
	TYPE_INT_BOOLEAN     = 0x12
	TYPE_FIRST_COLOR_INT = 0x1c
	TYPE_INT_COLOR_ARGB8 = 0x1c
	TYPE_INT_COLOR_RGB8  = 0x1d
	TYPE_INT_COLOR_ARGB4 = 0x1e
	TYPE_INT_COLOR_RGB4  = 0x1f
	TYPE_LAST_COLOR_INT  = 0x1f
	TYPE_LAST_INT        = 0x1f

	COMPLEX_UNIT_PX              = 0
	COMPLEX_UNIT_DIP             = 1
	COMPLEX_UNIT_SP              = 2
	COMPLEX_UNIT_PT              = 3
	COMPLEX_UNIT_IN              = 4
	COMPLEX_UNIT_MM              = 5
	COMPLEX_UNIT_SHIFT           = 0
	COMPLEX_UNIT_MASK            = 15
	COMPLEX_UNIT_FRACTION        = 0
	COMPLEX_UNIT_FRACTION_PARENT = 1
	COMPLEX_RADIX_23p0           = 0
	COMPLEX_RADIX_16p7           = 1
	COMPLEX_RADIX_8p15           = 2
	COMPLEX_RADIX_0p23           = 3
	COMPLEX_RADIX_SHIFT          = 4
	COMPLEX_RADIX_MASK           = 3
	COMPLEX_MANTISSA_SHIFT       = 8
	COMPLEX_MANTISSA_MASK        = 0xFFFFFF
)

type ResValue struct {
	/*
		res value header size
	*/
	Size uint16
	Res0 uint8
	/*
	 data type upside table
	*/
	DataType uint8
	/*
		data index
	*/
	Data uint32
}
