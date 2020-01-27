package buffer

import (
	"time"
)

var (
	OpenBuffers []*Buffer
)

type Buffer struct {
	BufferData
	cursorList []*Cursor
	currentCursorId int
	
}

type BufferData struct {
	ModTime    time.Time
	isModified bool
	ReadOnly   bool
	Filename   string
	Contents *LineArray
}

