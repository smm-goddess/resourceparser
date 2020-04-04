package items


const (
	SPEC_PUBLIC = uint32(0x40000000)
)

type ResTableTypeSpec struct {
	ResChunkHeader
	/*
		resource id:including: animator, anim,color,drawable,layout,raw,string,xml. each type has an id
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
