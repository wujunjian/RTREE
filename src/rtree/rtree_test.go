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

func TestSlice(t *testing.T) {

	var a []int

	func(b *[]int) {
		*b = append(*b, 1, 2, 3, 4)
	}(&a)

	fmt.Println(a)

	func(b []int) {
		b[0] = 9
		b = append(b, 5, 6, 7, 8)
	}(a)

	fmt.Println(a)
}

func TestRTree(x *testing.T) {

	var t RTree
	t.Init(4, 0.4, RTREE_QUADRATIC, 2)

	//北京动物园
	t.InsertCoordinates([]LngLatPoint{
		{116.328135, 39.942811},
		{116.336246, 39.943371},
		{116.342726, 39.942811},
		{116.340880, 39.945970},
		{116.341953, 39.939159},
		{116.329937, 39.939422}})

	//故宫
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}})

	// for i := 0; i < t.root.dataLength(); i++ {
	// fmt.Println("tree data :", t.root.getData(i).getLow(), t.root.getData(i).getHigh())
	// }

	//景山
	result := t.SearchCoordinates([]LngLatPoint{
		{116.399245, 39.925601},
	})

	fmt.Println("景山 result len", len(result))
	for _, r := range result {
		fmt.Println(r.low.data)
	}

	// 故宫
	result = t.SearchCoordinates([]LngLatPoint{
		{116.397614, 39.918920},
	})

	fmt.Println("故宫 result len", len(result))
	for _, r := range result {
		fmt.Println(r.low.data, r.high.data)
	}
}
