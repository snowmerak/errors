package bufferedlist

import "sync"

const nodeLength = 4096

var nodePool = sync.Pool{
	New: func() interface{} {
		return &node{}
	},
}

func newNode() *node {
	return nodePool.Get().(*node)
}

func popNode(n *node) {
	n.reset()
	nodePool.Put(n)
}

type node struct {
	data      [nodeLength]byte
	lastIndex int
	next      *node
}

func (nd *node) reset() {
	nd.lastIndex = 0
	nd.next = nil
}

func (nd *node) Write(p []byte) (n int, err error) {
	remains := nodeLength - nd.lastIndex
	if remains == 0 {
		return 0, nil
	}
	if remains < len(p) {
		copy(nd.data[nd.lastIndex:], p[:remains])
		nd.lastIndex = nodeLength
		return remains, nil
	}
	copy(nd.data[nd.lastIndex:], p)
	nd.lastIndex += len(p)
	return len(p), nil
}
