package items

const (
	END = 0xFFFFFFFF
)

type ResStringPoolSpan struct {
	Name      ResStringPoolRef
	FirstChar uint32
	LastChar  uint32
}
