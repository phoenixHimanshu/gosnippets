package main

//Method can have any type  including struct and any user defined type
// type level int
type Node struct {
	next *Node
	data int
}

func (ll *Node) SetValue(v int) {
	ll.data = v

}

//Method with reciver
func (ll *Node) GetValue() int {
	return ll.data
}

// Constructor writing style in golang
func NewLLCtor() *Node {
	//empty node is created, small memory allocated.
	return new(Node)
}

func main() {
	node := NewLLCtor()
	node.SetValue(10)
	println(node.GetValue())
}
