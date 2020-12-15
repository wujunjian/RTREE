package rtree

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
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

	// t.BFSearch()
	// fmt.Println("***************************************")
	// fmt.Println("***************************************")
	// fmt.Println()

	// for i, rect := range t.Range() {
	// 	fmt.Println(i, rect.info.(string))
	// }

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
		fmt.Println(r.toString(), r.info.(string))
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

func TestMillion(x *testing.T) {

	var t RTree
	t.Init(4, 0.4, RTREE_QUADRATIC, 2)

	// src := rand.NewSource(time.Now().Unix())
	src := rand.NewSource(1)

	r := rand.New(src)
	itemnum := 100000
	begintime := time.Now()
	for i := 0; i < itemnum; i++ {
		lng := math.Mod(r.NormFloat64()*360, 180)
		lat := math.Mod(r.NormFloat64()*180, 90)

		// fmt.Printf("%07d\t[%11.6f,%11.6f]\n", i, lng, lat)

		t.InsertCoordinates([]LngLatPoint{
			{lng, lat}}, fmt.Sprintf("%07d", i))

	}
	endtime := time.Now()
	fmt.Println("num", itemnum, "root level", t.root.getLevel(), "create usetime", endtime.Sub(begintime))

	searchRange := []LngLatPoint{
		{50, 80},
		{60, 90},
	}

	searchRange1 := []LngLatPoint{
		{110, 10},
		{120, 20},
	}
	begintime = time.Now()
	result := t.SearchCoordinates(searchRange)
	endtime = time.Now()
	fmt.Println("search", CoordinatesToRectangle(searchRange, nil).toString())
	fmt.Println("result", len(result), "search usetime", endtime.Sub(begintime))
	fmt.Println()

	begintime = time.Now()
	result = t.SearchCoordinates(searchRange1)
	endtime = time.Now()
	fmt.Println("search", CoordinatesToRectangle(searchRange1, nil).toString())
	fmt.Println("result", len(result), "search usetime", endtime.Sub(begintime))
	for _, item := range result {
		fmt.Println("info:", item.info.(string))
	}
	fmt.Println()

	// t.BFSearch()
	fmt.Println("***************************************")
	fmt.Println("***************************************")
	fmt.Println()

	begintime = time.Now()
	t.RDelete(result)
	endtime = time.Now()
	newresult := t.SearchCoordinates(searchRange1)
	fmt.Println("need zero newresult len", len(newresult), "delete usetime", endtime.Sub(begintime))
	//再次插入
	begintime = time.Now()
	for _, item := range result {
		fmt.Println("insert info", item.info.(string))
		t.Insert(item)
	}
	endtime = time.Now()
	fmt.Println("insert usetime", endtime.Sub(begintime))

	newresult1 := t.SearchCoordinates(searchRange1)
	fmt.Println("newresult1 len", len(newresult1))

}
