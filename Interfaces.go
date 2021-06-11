package main

type PowerNode struct {
	next *PowerNode
	val  int
}

type Inode interface {
	SetVal(i int)
	GetVal() int
}

func (ll *Node) SetVal(v int) {
	ll.data = v

}

func (ll *Node) GetVal() int {
	return ll.data
}

func NewCtor() *Node {
	return new(Node)
}

func (ll *PowerNode) SetVal(v int) {
	ll.val = v * 10
}
func (ll *PowerNode) GetVal() int {
	return ll.val
}
func NewPowerCtor() *PowerNode {
	return new(PowerNode)
}
