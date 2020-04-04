package items

type ResTableMapEntry struct {
	ResTableEntry
	Parent ResTableRef
	Count  uint32
}
