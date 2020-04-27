package buffer

import (
	"container/ring"
	"log"
	"strings"
	"time"

	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/linearray"
)

var (
	OpenBuffers []*Buffer
)

type Buffer struct {
	mode           interfaces.EditorMode
	data           *BufferData
	point          linearray.Loc
	mark           linearray.Loc
	scrollPosition linearray.Loc
	linesInDisplay int
	undoHistory    *ring.Ring
	redoHistory    *ring.Ring
}

type BufferData struct {
	modTime    time.Time
	isModified bool
	filename   string
	contents   *linearray.LineArray
}

type bufferSnapshot struct {
	timestamp        time.Time
	contents         *linearray.LineArray
	isModified       bool
	point            linearray.Loc
	mark             linearray.Loc
	scrollPosition   linearray.Loc
	modifiesContents bool
}

func (b *Buffer) GetLines() *linearray.LineArray {
	return b.data.contents
}

func (b *Buffer) LinesInDisplay() int {
	return b.linesInDisplay
}
func (b *Buffer) SetLinesInDisplay(v int) {
	b.linesInDisplay = v
}

func (b *Buffer) Mark() linearray.Loc {
	return b.mark
}
func (b *Buffer) Point() linearray.Loc {
	return b.point
}

func (b *Buffer) SetMark(l linearray.Loc) {
	b.mark = l
}

func (b *Buffer) SetPoint(l linearray.Loc) {
	b.point = l
}

func (b *Buffer) Mode() interfaces.EditorMode {
	return b.mode
}

func (b *Buffer) ScrollPosition() int {
	return b.scrollPosition.Y
}

func (b *Buffer) SetScrollPosition(v int) {
	b.scrollPosition.Y = v
}

func (b *Buffer) Undo() {
	snap, ok := b.undoHistory.Value.(*bufferSnapshot)
	if ok && snap != nil {
		b.undoHistory = b.undoHistory.Prev()
		currentState := b.snapshot(snap.ModifiesBuffer())
		b.RestoreSnapshot(snap)
		b.redoHistory = b.redoHistory.Next()
		b.redoHistory.Value = currentState
	}
}

func (b *Buffer) Redo() {
	snap, ok := b.redoHistory.Value.(*bufferSnapshot)
	if ok && snap != nil {
		b.redoHistory = b.redoHistory.Prev()
		currentState := b.snapshot(snap.ModifiesBuffer())
		b.RestoreSnapshot(snap)
		b.undoHistory = b.undoHistory.Next()
		b.undoHistory.Value = currentState
	}
}

func (b *Buffer) snapshot(willChangeContents bool) *bufferSnapshot {
	snap := &bufferSnapshot{}
	snap.timestamp = time.Now()
	snap.isModified = b.data.isModified
	snap.point = b.point
	snap.mark = b.mark
	snap.scrollPosition = b.scrollPosition
	if willChangeContents {
		snap.contents = b.data.contents.Copy()
	} else {
		snap.contents = nil
	}

	return snap
}

func (b *Buffer) TakeSnapshot(willChangeContents bool) {
	// clear the redo history
	b.redoHistory = newRing()
	snap := b.snapshot(willChangeContents)
	// log.Printf("Snapbuf %v %v\n", snap.timestamp, string(snap.contents.Bytes()))
	b.undoHistory = b.undoHistory.Next()
	b.undoHistory.Value = snap
}

func (b *Buffer) RestoreSnapshot(snap *bufferSnapshot) {
	b.data.isModified = snap.isModified
	b.point = snap.point
	b.mark = snap.mark
	b.scrollPosition = snap.scrollPosition
	if snap.contents != nil {
		b.data.contents = snap.contents
		// log.Printf("Restore %v %v\n", snap.timestamp, string(snap.contents.Bytes()))
	}
}

func (snap *bufferSnapshot) ModifiesBuffer() bool {
	return snap.contents != nil
}

type myBufferFactory struct{}

func (m myBufferFactory) NewScratchBuffer() interfaces.Buffer {

	sb := newScratchBuffer()

	OpenBuffers = append(OpenBuffers, sb)
	log.Printf("OpenBuffers: %v\n", OpenBuffers)
	return sb
}

func (m myBufferFactory) CreateBuffer(mode string) interfaces.Buffer {

	sb := newBuffer(mode)
	OpenBuffers = append(OpenBuffers, sb)
	log.Printf("OpenBuffers: %v\n", OpenBuffers)
	return sb
}

func newScratchBuffer() *Buffer {
	la := linearray.NewLineArray(100, strings.NewReader("*scratch*\nhello world\nthis is a temp buffer\n"))
	ub := newRing()
	rb := newRing()
	bd := &BufferData{time.Now(), false, "*scratch*", la}
	m := interfaces.GetMode("edit")
	b := &Buffer{m, bd, linearray.Loc{0, 0}, linearray.Loc{-1, -1}, linearray.Loc{0, 0}, 1, ub, rb}
	log.Printf("New scratch buffer: %v", b)
	return b
}

func newBuffer(mode string) *Buffer {
	la := linearray.NewLineArray(100, strings.NewReader("\n"))
	ub := newRing()
	rb := newRing()
	bd := &BufferData{time.Now(), false, "", la}
	m := interfaces.GetMode(mode)
	b := &Buffer{m, bd, linearray.Loc{0, 0}, linearray.Loc{-1, -1}, linearray.Loc{0, 0}, 1, ub, rb}
	return b
}

func newRing() *ring.Ring {
	return ring.New(100)
}
