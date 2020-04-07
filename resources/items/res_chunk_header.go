package items

type ResChunkHeader struct {
	/*
	 chunk type
	*/
	Type uint16
	/*
		chunk header size
	*/
	HeaderSize uint16
	/*
		chunk size
	*/
	Size uint32
}

//func (chunkHeader ResChunkHeader) String() string {
//	return fmt.Sprintf("header type:%x, header size:%d, chunk size: %d", chunkHeader.Type, chunkHeader.HeaderSize, chunkHeader.Size)
//}

//func (this ResChunkHeader) GetSize() uint32 {
//	return 8
//}
