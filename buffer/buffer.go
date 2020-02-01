package buffer

import (
	"errors"
	"strings"
	"time"
)

var (
	OpenBuffers []*Buffer
)

type Buffer struct {
	Data  *BufferData
	Point *Loc
	Mark  *Loc
}

type BufferData struct {
	ModTime    time.Time
	isModified bool
	ReadOnly   bool
	Filename   string
	Contents   *LineArray
}

func GetScratchBuffer() *Buffer {
	if sb, err := findScratchBuffer(); err != nil {
		return sb
	}

	sb := newScratchBuffer()
	OpenBuffers = append(OpenBuffers, sb)
	return sb
}

func findScratchBuffer() (*Buffer, error) {
	for _, b := range OpenBuffers {
		if b.Data.Filename == "*scratch*" {
			return b, nil
		}
	}

	return nil, errors.New("scratch not found")
}

func newScratchBuffer() *Buffer {
	la := NewLineArray(100, strings.NewReader("*scratch*\nhello world\n"))
	bd := &BufferData{time.Now(), false, false, "*scratch*", la}
	b := &Buffer{bd, &Loc{0, 0}, nil}
	return b
}
