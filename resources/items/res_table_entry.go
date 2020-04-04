package items

const (
	FLAG_COMPLEX = uint32(0x0001)
	FLAG_PUBLIC  = uint32(0x0002)
)

type ResTableEntry struct {
	Size uint16
	/*
		flags = 1, data type is ResTableMapEntry, flags = 0 data type is ResValue
	*/
	Flags uint16
	Key   ResStringPoolRef
}
