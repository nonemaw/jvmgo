package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}
