package bufferedlist

import "io"

type BufferedList struct {
	head *node
	tail *node
}

func New() *BufferedList {
	return &BufferedList{}
}

func (b *BufferedList) Write(p []byte) (n int, err error) {
	if b.head == nil {
		b.head = newNode()
		b.tail = b.head
	}

	length := len(p)
	for length > 0 {
		written, err := b.tail.Write(p)
		if err != nil {
			return n, err
		}
		length -= written
		n += written
		if length > 0 && written == 0 {
			b.tail.next = newNode()
			b.tail = b.tail.next
		}
	}

	return n, nil
}

func (b *BufferedList) WriteByte(c byte) error {
	n, err := b.Write([]byte{c})
	if err != nil {
		return err
	}

	if n != 1 {
		return io.ErrShortBuffer
	}

	return nil
}

func (b *BufferedList) WriteString(s string) (n int, err error) {
	n, err = b.Write([]byte(s))
	if err != nil {
		return n, err
	}

	if n != len(s) {
		return n, io.ErrShortBuffer
	}

	return n, nil
}

func (b *BufferedList) Length() int64 {
	if b.head == nil {
		return 0
	}

	length := int64(0)

	for n := b.head; n != nil; n = n.next {
		length += int64(n.lastIndex)
	}

	return length
}

func (b *BufferedList) Free() {
	for n := b.head; n != nil; n = n.next {
		popNode(n)
	}
	b.head = nil
	b.tail = nil
}
