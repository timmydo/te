package buffer

import (
	"log"
)

// Loc stores a location
type Loc struct {
	X, Y int
}

// LessThan returns true if b is smaller
func (l Loc) LessThan(b Loc) bool {
	if l.Y < b.Y {
		return true
	}
	return l.Y == b.Y && l.X < b.X
}

// GreaterThan returns true if b is bigger
func (l Loc) GreaterThan(b Loc) bool {
	if l.Y > b.Y {
		return true
	}
	return l.Y == b.Y && l.X > b.X
}

// GreaterEqual returns true if b is greater than or equal to b
func (l Loc) GreaterEqual(b Loc) bool {
	if l.Y > b.Y {
		return true
	}
	if l.Y == b.Y && l.X > b.X {
		return true
	}
	return l == b
}

// LessEqual returns true if b is less than or equal to b
func (l Loc) LessEqual(b Loc) bool {
	if l.Y < b.Y {
		return true
	}
	if l.Y == b.Y && l.X < b.X {
		return true
	}
	return l == b
}

// The following functions require a buffer to know where newlines are

// Diff returns the distance between two locations
func DiffLA(a, b Loc, buf *LineArray) int {
	if a.Y == b.Y {
		if a.X > b.X {
			return a.X - b.X
		}
		return b.X - a.X
	}

	// Make sure a is guaranteed to be less than b
	if b.LessThan(a) {
		a, b = b, a
	}

	loc := 0
	for i := a.Y + 1; i < b.Y; i++ {
		// + 1 for the newline
		loc += buf.RuneCount(i) + 1
	}
	loc += buf.RuneCount(a.Y) - a.X + b.X + 1
	return loc
}

// This moves the location one character to the right
func (l Loc) right(buf *LineArray) Loc {
	if l == buf.End() {
		return Loc{l.X + 1, l.Y}
	}
	var res Loc
	if l.X < buf.RuneCount(l.Y) {
		res = Loc{l.X + 1, l.Y}
	} else {
		res = Loc{0, l.Y + 1}
	}
	return res
}

// This moves the given location one character to the left
func (l Loc) left(buf *LineArray) Loc {
	if l == buf.Start() {
		return Loc{l.X - 1, l.Y}
	}
	var res Loc
	if l.X > 0 {
		res = Loc{l.X - 1, l.Y}
	} else {
		res = Loc{buf.RuneCount(l.Y - 1), l.Y - 1}
	}
	return res
}

// Move moves the cursor n characters to the left or right
// It moves the cursor left if n is negative
func (l Loc) MoveLA(n int, buf *LineArray) Loc {
	if n > 0 {
		for i := 0; i < n; i++ {
			l = l.right(buf)
		}
		return l
	}
	for i := 0; i < -n; i++ {
		l = l.left(buf)
	}

	lineRuneCount := buf.RuneCount(l.Y)
	if l.X > lineRuneCount {
		l = Loc{lineRuneCount, l.Y}
	}

	return l
}

func (l Loc) Diff(a, b Loc, buf *Buffer) int {
	return DiffLA(a, b, buf.Data.Contents)
}
func (l Loc) Move(n int, buf *Buffer) Loc {
	return l.MoveLA(n, buf.Data.Contents)
}

func (l Loc) MoveDownLines(n int, buf *Buffer) Loc {
	log.Printf("End: %v\n", buf.Data.Contents.End())
	if l.Y+n < 0 {
		return Loc{l.X, 0}
	}

	if l.Y+n > buf.Data.Contents.End().Y {
		return Loc{l.X, buf.Data.Contents.End().Y}
	}

	return Loc{l.X, l.Y + n}
}

func (l Loc) MoveInBounds(n int, buf *Buffer) Loc {
	newPos := l.Move(n, buf)
	if newPos.GreaterEqual(buf.Data.Contents.End()) {
		return buf.Data.Contents.End()
	}

	if newPos.LessEqual(buf.Data.Contents.Start()) {
		return buf.Data.Contents.Start()
	}

	return newPos
}

// ByteOffset is just like ToCharPos except it counts bytes instead of runes
func ByteOffset(pos Loc, buf *Buffer) int {
	x, y := pos.X, pos.Y
	loc := 0
	for i := 0; i < y; i++ {
		// + 1 for the newline
		loc += len(buf.Data.Contents.LineBytes(i)) + 1
	}
	loc += len(buf.Data.Contents.LineBytes(y)[:x])
	return loc
}
