package rtree

import "math"

// RTNode ...
type RTNode struct {
	rtree       *RTree      // 结点所在的树
	level       int         // 结点所在的层
	datas       []Rectangle // 相当于条目
	parent      *RTNode     // 父节点
	usedSpace   int         // 结点已用的空间
	insertIndex int         // 记录插入的搜索路径索引
	deleteIndex int         // 记录删除的查找路径索引
}

func (r *RTNode) init(rtree *RTree, paraent *RTNode, level int) {
	r.rtree = rtree
	r.parent = paraent
	r.level = level

	r.datas = make([]Rectangle, r.rtree.getNodeCapacity()+1) // 多出的一个用于节点分裂
	r.usedSpace = 0
}

func (r RTNode) getParent() *RTNode {
	return r.parent
}

func (r RTNode) isRoot() bool {
	return r.parent == nil
}

// 是否非叶子节点
func (r RTNode) isIndex() bool {
	return r.level != 0
}

// 是否叶子节点
func (r RTNode) isLeaf() bool {
	return r.level == 0
}

func (r *RTNode) addData(rect Rectangle) {
	if r.usedSpace == r.rtree.getNodeCapacity() {
		panic(1) //Node is full.
	}

	r.datas[r.usedSpace] = rect
	r.usedSpace++

}

func (r *RTNode) deleteData(i int) {
	if !r.datas[i+1].isNULL() {
		var tmp []Rectangle
		tmp = append(tmp, r.datas[0:i]...)
		tmp = append(tmp, r.datas[i+1:]...)
		r.datas = tmp
	} else {
		r.datas[i].clean()
	}

	r.usedSpace--
}

// 压缩算法 叶节点L中刚刚删除了一个条目，
// 如果这个结点的条目数太少下溢， 则删除该结点，同时将该结点中剩余的条目重定位到其他结点中。
// 如果有必要，要逐级向上进行这种删除，调整向上传递的路径上的所有外包矩形，使其尽可能小，直到根节点。
// 存储删除结点中剩余条目
func (r *RTNode) condenseTree(c []RTNode) {

}

// --> 插入新的Rectangle后从插入的叶节点开始向上调整RTree，直到根节点
// node1 引起需要调整的孩子结点
// node2 分裂的结点，若未分裂则为null
func (r *RTNode) adjustTree(node1 *RTNode, node2 *RTNode) {

}

//
// 分裂结点的平方算法
// 1、为两个组选择第一个条目--调用算法pickSeeds()来为两个组选择第一个元素，分别把选中的两个条目分配到两个组当中。<br>
// 2、检查是否已经分配完毕，如果一个组中的条目太少，为避免下溢，将剩余的所有条目全部分配到这个组中，算法终止<br>
// 3、调用pickNext来选择下一个进行分配的条目--计算把每个条目加入每个组之后面积的增量，选择两个组面积增量差最大的条目索引,
// 如果面积增量相等则选择面积较小的组，若面积也相等则选择条目数更少的组<br>
//
// @param rectangle
//            导致分裂的溢出Rectangle
// @return 两个组中的条目的索引
func (r *RTNode) quadraticSplit(rect Rectangle) [][]int {
	r.datas[r.usedSpace] = rect //先添加进去
	total := r.usedSpace + 1    // 结点总数

	mask := make([]int, total) // 标记访问的条目

	for i := 0; i < total; i++ {
		mask[i] = 1
	}

	// 分裂后每个组只是有total/2个条目
	c := total/2 + 1
	minNodeSize := int(math.Round(float64(r.rtree.getNodeCapacity()) * r.rtree.fillFactor))
	if minNodeSize < 2 {
		minNodeSize = 2
	}
	// 记录没有被检查的条目的个数
	rem := total

	group1 := make([]int, c)
	group2 := make([]int, c)

	// 跟踪被插入每个组的条目的索引
	var i1, i2 int

	seed := r.pickSeeds()
	group1[i1] = seed[0]
	mask[seed[0]] = -1
	i1++

	group2[i2] = seed[1]
	mask[seed[1]] = -1
	i2++

	rem -= 2

	for rem > 0 {
		// 将剩余的所有条目全部分配到group1组中，算法终止
		if minNodeSize-i1 == rem {

			// 将剩余的所有条目全部分配到group1组中，算法终止
		} else if minNodeSize-i2 == rem {

		} else {

		}
	} // end while

	ret := make([][]int, 2)
	ret[0] = group1
	ret[1] = group2

	return ret
}

// 1、对每一对条目E1和E2，计算包围它们的Rectangle J，
//    计算d = area(J) - area(E1) - area(E2);
// 2、Choose the pair with the largest d
// @return 返回两个条目如果放在一起会有最多的冗余空间的条目索引
func (r RTNode) pickSeeds() []int {
	var inefficiency float64
	var i1, i2 int

	for i := 0; i < r.usedSpace; i++ {
		for j := i + 1; j <= r.usedSpace; j++ {
			rect := r.datas[i].getUnionRectangle(r.datas[j])
			d := rect.getArea() - r.datas[i].getArea() - r.datas[j].getArea()

			if d > inefficiency {
				inefficiency = d
				i1 = i
				i2 = j
			}

		}
	}
	return []int{i1, i2}
}
