package interfaces

import "github.com/timmydo/te/linearray"

type Buffer interface {
	Mode() EditorMode
	GetLines() *linearray.LineArray
	Point() linearray.Loc
	Mark() linearray.Loc
	SetPoint(linearray.Loc)
	SetMark(linearray.Loc)
	ScrollPosition() int
	SetScrollPosition(int)
	LinesInDisplay() int
	SetLinesInDisplay(int)
	TakeSnapshot(bool)
	Undo()
	Redo()
}

type BufferFactory interface {
	NewScratchBuffer() Buffer
	CreateBuffer(string) Buffer
	CreateBufferFromFile(string) (Buffer, error)
	DeleteBuffer(Buffer) (bool, Buffer)
}

var bufferFactory BufferFactory

func GetBufferFactory() BufferFactory {
	return bufferFactory
}

func SetBufferFactory(bf BufferFactory) {
	bufferFactory = bf
}
