package main

import (
	"fmt"
	"testing"
)

func Test_Method(t *testing.T) {
	node := NewLLCtor()
	node.SetValue(10)
	println(node.GetValue())

}

func Test_Interface(t *testing.T) {

	//why node is not a pointer to Inode
	var node Inode
	node = NewCtor()
	node.SetVal(40)
	println(node.GetVal())

	//Now node is pointing to PowerNode (polymorphism)
	node = NewPowerCtor()
	node.SetVal(10)
	println(node.GetVal())

	//How to get value of concrete type from interface variable
	if n, ok := node.(*PowerNode); ok {
		fmt.Println(n.val)

	}

}

func Test_runGoRoutine(t *testing.T) {
	runGoRoutine()
}
func Test_ChInAction(t *testing.T) {
	ChInAction()
}
