package rtree

import (
	"fmt"
	"testing"
)

type myin interface {
	myfunc()
}

type father struct {
}

func (f father) myfunc() {
	fmt.Println("father")
}

type child struct {
	father
}

func (f child) myfunc() {
	fmt.Println("child")
}

func TestMyinterface(t *testing.T) {
	var c child
	c.myfunc()
	c.father.myfunc()

	var f father
	f.myfunc()

	var myini myin
	myini = c
	myini.myfunc()

	myini = f
	myini.myfunc()
}

func TestRTree(t *testing.T) {

	var a = &RTree{}

	fmt.Println("aaa", a)

}
