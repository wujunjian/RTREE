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
		{116.329937, 39.939422}}, "zoo")

	//故宫
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫0")
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫1")
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫2")
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫3")
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫4")
	t.InsertCoordinates([]LngLatPoint{
		{116.392035, 39.922475},
		{116.401777, 39.922902},
		{116.402078, 39.913423},
		{116.392336, 39.913193}}, "故宫5")

	//家
	t.InsertCoordinates([]LngLatPoint{
		{116.556916, 39.919734},
		{116.559792, 39.919784},
		{116.559867, 39.918015},
		{116.556906, 39.917784}}, "home")

	for i, rect := range t.Range() {
		fmt.Println(i, rect.info.(string))
	}
	return

	//景山
	result := t.SearchCoordinates([]LngLatPoint{
		{116.399245, 39.925601},
	})

	fmt.Println("景山 result len", len(result))
	for _, r := range result {
		fmt.Println(r.low.data, r.high.data)
	}

	// 故宫
	result = t.SearchCoordinates([]LngLatPoint{
		{116.397614, 39.918920},
	})

	fmt.Println("result len", len(result))
	for _, r := range result {
		fmt.Println(r.low.data, r.high.data, r.info.(string))
	}
}

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a[0:2])
	fmt.Println(b, len(b), cap(b))
	copy(b[2:], a[3:])
	fmt.Println(b, len(b), cap(b))
}
