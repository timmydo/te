package buffer

import (
	"container/ring"
	"log"
	"strings"
	"time"

	"github.com/timmydo/te/theme"
)

var (
	OpenBuffers []*Buffer
)

type Buffer struct {
	Mode           string
	Data           *BufferData
	Point          Loc
	Mark           Loc
	ScrollPosition Loc
	LinesInDisplay int
	UndoHistory    *ring.Ring
	RedoHistory    *ring.Ring
	StyleProvider  BufferStyleProvider
}

type BufferStyleProvider interface {
	GetBufferStyle() *theme.BufferThemeStyle
	GetLineStyle(int) *theme.LineThemeStyle
	GetCharacterStyle(int, int) *theme.CharacterThemeStyle
}

type DefaultStyleProvider struct{}

func (this *DefaultStyleProvider) GetBufferStyle() *theme.BufferThemeStyle {
	return theme.DefaultBufferTheme
}

func (this *DefaultStyleProvider) GetLineStyle(int) *theme.LineThemeStyle {
	return theme.DefaultLineTheme
}

func (this *DefaultStyleProvider) GetCharacterStyle(int, int) *theme.CharacterThemeStyle {
	return theme.DefaultCharacterTheme
}

type BufferData struct {
	ModTime    time.Time
	isModified bool
	Filename   string
	Contents   *LineArray
}

type BufferSnapshot struct {
	timestamp        time.Time
	contents         *LineArray
	isModified       bool
	point            Loc
	mark             Loc
	scrollPosition   Loc
	ModifiesContents bool
}

func (b *Buffer) Snapshot(willChangeContents bool) *BufferSnapshot {
	snap := &BufferSnapshot{}
	snap.timestamp = time.Now()
	snap.isModified = b.Data.isModified
	snap.point = b.Point
	snap.mark = b.Mark
	snap.scrollPosition = b.ScrollPosition
	if willChangeContents {
		snap.contents = b.Data.Contents.Copy()
	} else {
		snap.contents = nil
	}

	return snap
}

func (b *Buffer) TakeSnapshot(willChangeContents bool) {
	// clear the redo history
	b.RedoHistory = newRing()
	snap := b.Snapshot(willChangeContents)
	// log.Printf("Snapbuf %v %v\n", snap.timestamp, string(snap.contents.Bytes()))
	b.UndoHistory = b.UndoHistory.Next()
	b.UndoHistory.Value = snap
}

func (b *Buffer) RestoreSnapshot(snap *BufferSnapshot) {
	b.Data.isModified = snap.isModified
	b.Point = snap.point
	b.Mark = snap.mark
	b.ScrollPosition = snap.scrollPosition
	if snap.contents != nil {
		b.Data.Contents = snap.contents
		// log.Printf("Restore %v %v\n", snap.timestamp, string(snap.contents.Bytes()))
	}
}

func (snap *BufferSnapshot) ModifiesBuffer() bool {
	return snap.contents != nil
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
	ub := newRing()
	rb := newRing()
	bd := &BufferData{time.Now(), false, "*scratch*", la}
	b := &Buffer{"edit", bd, Loc{0, 0}, Loc{-1, -1}, Loc{0, 0}, 1, ub, rb, &DefaultStyleProvider{}}
	log.Printf("New scratch buffer: %v", b)
	return b
}

func newRing() *ring.Ring {
	return ring.New(100)
}
