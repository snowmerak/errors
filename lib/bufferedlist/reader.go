package bufferedlist

import (
	"io"
)

type Reader struct {
	head *node
	tail *node

	currentNode *node
	nodeCursor  int64
	byteCursor  int64
}

func (b *BufferedList) Reader() *Reader {
	reader := &Reader{
		head: newNode(),
		tail: nil,
	}
	reader.tail = reader.head

	for n := b.head; n != nil; n = n.next {
		copy(reader.tail.data[:], n.data[:])
		reader.tail.lastIndex = n.lastIndex
		if n.next != nil {
			reader.tail.next = newNode()
		}
	}

	return reader
}

func (r *Reader) Seek(offset int64, _ int) (int64, error) {
	nodeCursor := offset / nodeLength
	byteCursor := offset % nodeLength

	for i := int64(0); i < nodeCursor; i++ {
		if r.currentNode.next == nil {
			return -1, io.EOF
		}
		r.currentNode = r.currentNode.next
	}

	r.nodeCursor = nodeCursor
	r.byteCursor = byteCursor

	return offset, nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.currentNode == nil {
		r.currentNode = r.head
	}

	length := len(p)
	for length > 0 {
		if r.currentNode == nil {
			return n, io.EOF
		}

		remains := r.currentNode.lastIndex - int(r.byteCursor)
		if remains == 0 {
			return n, io.EOF
		}

		if remains > length {
			remains = length
		}

		copy(p[n:], r.currentNode.data[r.byteCursor:r.byteCursor+int64(remains)])
		n += remains
		length -= remains
		r.byteCursor += int64(remains)
		if r.byteCursor == nodeLength {
			r.currentNode = r.currentNode.next
			r.byteCursor = 0
			r.nodeCursor++
		}
	}

	return n, nil
}

func (r *Reader) Free() {
	for n := r.head; n != nil; {
		next := n.next
		popNode(n)
		n = next
	}
	r.head = nil
	r.tail = nil
}
