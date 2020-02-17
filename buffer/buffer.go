package buffer

import (
	"log"
	"strings"
	"time"
)

var (
	OpenBuffers []*Buffer
)

type Buffer struct {
	Data           *BufferData
	Point          Loc
	Mark           Loc
	ScrollPosition Loc
	LinesInDisplay int
}

type BufferData struct {
	ModTime    time.Time
	isModified bool
	ReadOnly   bool
	Filename   string
	Contents   *LineArray
}

func GetScratchBuffer() *Buffer {
	if sb := findScratchBuffer(); sb != nil {
		return sb
	}

	sb := newScratchBuffer()

	OpenBuffers = append(OpenBuffers, sb)
	log.Printf("OpenBuffers: %v\n", OpenBuffers)
	return sb
}

func findScratchBuffer() *Buffer {
	for _, b := range OpenBuffers {
		if b.Data.Filename == "*scratch*" {
			return b
		}
	}

	return nil
}

func newScratchBuffer() *Buffer {
	la := NewLineArray(100, strings.NewReader("*scratch*\nhello world\nthis is a temp buffer\n"))
	bd := &BufferData{time.Now(), false, false, "*scratch*", la}
	b := &Buffer{bd, Loc{0, 0}, Loc{-1, -1}, Loc{0, 0}, 1}
	log.Printf("New scratch buffer: %v", b)
	return b
}
