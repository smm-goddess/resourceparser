package items

const (
	NO_ENTRY = 0xFFFFFFFF
)

/*
	follow ResTableTypeSpec is a uint32 array, size equals entry count, each item is the diff of the item
*/

type ResTableType struct {
	ResChunkHeader
	/*
		type id
	*/
	Id   uint8
	Res0 uint8
	Res  uint16
	/*
		resource count
	*/
	EntryCount uint32
	/*
		resource data offset relative to header
	*/
	EntriesStart uint32
	/*
		ResTableConfig, locale, language, resolution etc.
	*/
	ResTableConfig
}
