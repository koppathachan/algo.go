package bst

type Node interface {
	IsLess(node Node) bool
}

type node struct {
	left  *node
	right *node
	val   Node
}

func insert(n *node, val Node) {
	if n == nil {
		return
	}
	if val.IsLess(n.val) {
		if n.right == nil {
			n.right = &node{val: val}
		} else {
			insert(n.right, val)
		}
	} else {
		if n.left == nil {
			n.left = &node{val: val}
		} else {
			insert(n.left, val)
		}
	}
}

func (n *node) Insert(val Node) {
	if n.val == nil {
		n.val = val
	} else {
		insert(n, val)
	}
}

func inOrder(n *node) []Node {
	if n == nil {
		return []Node{}
	}
	return append(append(inOrder(n.left), n.val), inOrder(n.right)...)
}

func postOrder(n *node) []Node {
	if n == nil {
		return []Node{}
	}
	return append(append(postOrder(n.right), postOrder(n.left)...), n.val)
}

func preOrder(n *node) []Node {
	if n == nil {
		return []Node{}
	}
	return append([]Node{n.val}, append(preOrder(n.left), preOrder(n.right)...)...)
}

func (n *node) InOrder() []Node {
	return inOrder(n)
}
func (n *node) PreOrder() []Node {
	return preOrder(n)
}
func (n *node) PostOrder() []Node {
	return postOrder(n)
}

type Root interface {
	Insert(node Node)
	InOrder() []Node
	PreOrder() []Node
	PostOrder() []Node
}

func NewRoot() Root {
	return &node{}
}
