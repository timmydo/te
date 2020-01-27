package buffer


type Cursor struct {
	buf *Buffer
	Point Loc

	// where selection starts
	Mark Loc

	Id int
}

func NewCursor(b *Buffer, l Loc) *Cursor {
	c := &Cursor{
		buf: b,
		Loc: l,
	}

	return c
}

func (c *Cursor) Buf() *Buffer {
	return c.buf
}
