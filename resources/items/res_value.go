package items

const (
	// Contains no data.
	TYPE_NULL = 0x00
	// The 'data' holds a ResTable_ref, a reference to another resource
	// table entry.
	TYPE_REFERENCE = 0x01
	// The 'data' holds an attribute resource identifier.
	TYPE_ATTRIBUTE = 0x02
	// The 'data' holds an index into the containing resource table's
	// global value string pool.
	TYPE_STRING = 0x03
	// The 'data' holds a single-precision floating point number.
	TYPE_FLOAT = 0x04
	// The 'data' holds a complex number encoding a dimension value,
	// such as "100in".
	TYPE_DIMENSION = 0x05
	// The 'data' holds a complex number encoding a fraction of a
	// container.
	TYPE_FRACTION = 0x06
	// Beginning of integer flavors...
	TYPE_FIRST_INT = 0x10
	// The 'data' is a raw integer value of the form n..n.
	TYPE_INT_DEC = 0x10
	// The 'data' is a raw integer value of the form 0xn..n.
	TYPE_INT_HEX = 0x11
	// The 'data' is either 0 or 1, for input "false" or "true" respectively.
	TYPE_INT_BOOLEAN = 0x12
	// Beginning of color integer flavors...
	TYPE_FIRST_COLOR_INT = 0x1c
	// The 'data' is a raw integer value of the form #aarrggbb.
	TYPE_INT_COLOR_ARGB8 = 0x1c
	// The 'data' is a raw integer value of the form #rrggbb.
	TYPE_INT_COLOR_RGB8 = 0x1d
	// The 'data' is a raw integer value of the form #argb.
	TYPE_INT_COLOR_ARGB4 = 0x1e
	// The 'data' is a raw integer value of the form #rgb.
	TYPE_INT_COLOR_RGB4 = 0x1f
	// ...end of integer flavors.
	TYPE_LAST_COLOR_INT = 0x1f
	// ...end of integer flavors.
	TYPE_LAST_INT = 0x1f
)
const (
	// Where the unit type information is.  This gives us 16 possible
	// types as defined below.
	COMPLEX_UNIT_SHIFT = 0
	COMPLEX_UNIT_MASK  = 0xf
	// TYPE_DIMENSION: Value is raw pixels.
	COMPLEX_UNIT_PX = 0
	// TYPE_DIMENSION: Value is Device Independent Pixels.
	COMPLEX_UNIT_DIP = 1
	// TYPE_DIMENSION: Value is a Scaled device independent Pixels.
	COMPLEX_UNIT_SP = 2
	// TYPE_DIMENSION: Value is in points.
	COMPLEX_UNIT_PT = 3
	// TYPE_DIMENSION: Value is in inches.
	COMPLEX_UNIT_IN = 4
	// TYPE_DIMENSION: Value is in millimeters.
	COMPLEX_UNIT_MM = 5
	// TYPE_FRACTION: A basic fraction of the overall size.
	COMPLEX_UNIT_FRACTION = 0
	// TYPE_FRACTION: A fraction of the parent size.
	COMPLEX_UNIT_FRACTION_PARENT = 1
	// Where the radix information is, telling where the decimal place
	// appears in the mantissa.  This give us 4 possible fixed point
	// representations as defined below.
	COMPLEX_RADIX_SHIFT = 4
	COMPLEX_RADIX_MASK  = 0x3
	// The mantissa is an integral number -- i.e. 0xnnnnnn.0
	COMPLEX_RADIX_23p0 = 0
	// The mantissa magnitude is 16 bits -- i.e. 0xnnnn.nn
	COMPLEX_RADIX_16p7 = 1
	// The mantissa magnitude is 8 bits -- i.e. 0xnn.nnnn
	COMPLEX_RADIX_8p15 = 2
	// The mantissa magnitude is 0 bits -- i.e. 0x0.nnnnnn
	COMPLEX_RADIX_0p23 = 3
	// Where the actual value is.  This gives us 23 bits of
	// precision.  The top bit is the sign.
	COMPLEX_MANTISSA_SHIFT = 8
	COMPLEX_MANTISSA_MASK  = 0xffffff
)

type TYPE uint32

func (t TYPE) GetTypeDescription() string {
	switch t {
	case TYPE_NULL:
		return "null"
	case TYPE_REFERENCE:
		return "reference"
	case TYPE_ATTRIBUTE:
		return "attribute"
	case TYPE_STRING:
		return "string"
	case TYPE_FLOAT:
		return "float"
	case TYPE_DIMENSION:
		return "dimension"
	case TYPE_FRACTION:
		return "fraction"
	case TYPE_FIRST_INT:
		return "first int"
	//case TYPE_INT_DEC:
	case TYPE_INT_HEX:
		return "int hex"
	case TYPE_INT_BOOLEAN:
		return "boolean"
	case TYPE_FIRST_COLOR_INT:
		return "first color int"
	//case TYPE_INT_COLOR_ARGB8:
	case TYPE_INT_COLOR_RGB8:
		return "int color rgb8"
	case TYPE_INT_COLOR_ARGB4:
		return "int color argb4"
	//case TYPE_INT_COLOR_RGB4:
	case TYPE_LAST_COLOR_INT:
		return "last color int"
		//case TYPE_LAST_INT:
	default:
		return "unknown"
	}
}

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
