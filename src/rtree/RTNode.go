package rtree

import (
	"fmt"
	"math"
)

type IRTNode interface {
	getParent() IRTNode
	getChild(int) IRTNode
	delChild(int)
	setParent(IRTNode)
	isRoot() bool
	isIndex() bool
	isLeaf() bool
	addData(Rectangle)
	getData(int) Rectangle
	deleteData(int)
	setDatas(int, Rectangle)
	getDeleteIndex() int
	condenseTree(*[]IRTNode)
	adjustTree(IRTNode, IRTNode)
	quadraticSplit(Rectangle) [][]int
	pickSeeds() []int
	getNodeRectangle() Rectangle
	chooseLeaf(Rectangle) *RTDataNode
	findLeaf(Rectangle) *RTDataNode //用于删除

	Search(Rectangle) []Rectangle //用于查找

	getUsedSpace() int
	getLevel() int
}

// RTNode ...
type RTNode struct {
	rtree       *RTree      // 结点所在的树
	level       int         // 结点所在的层
	datas       []Rectangle // 相当于条目
	parent      IRTNode     // 父节点
	usedSpace   int         // 结点已用的空间
	insertIndex int         // 记录插入的搜索路径索引
	deleteIndex int         // 记录删除的查找路径索引
}

func (r *RTNode) init(rtree *RTree, paraent IRTNode, level int) {
	r.rtree = rtree
	r.parent = paraent
	r.level = level

	r.datas = make([]Rectangle, r.rtree.getNodeCapacity()+1) // 多出的一个用于节点分裂
	r.usedSpace = 0
}

func (r RTNode) getLevel() int {
	return r.level
}

func (r RTNode) getUsedSpace() int {
	return r.usedSpace
}

func (r RTNode) Search(Rectangle) []Rectangle {
	fmt.Println("RTNode Search will never be called")
	// panic("RTNode Search will never be called")
	return nil
}

func (r RTNode) getDeleteIndex() int {
	return r.deleteIndex
}

func (r RTNode) getChild(index int) IRTNode {
	defer panic("RTNode getChild will never be called")
	return nil
}
func (r *RTNode) delChild(index int) {
	panic("RTNode delChild will never be called")
}

func (r RTNode) getParent() IRTNode {
	return r.parent
}

func (r *RTNode) setParent(node IRTNode) {
	r.parent = node
}

func (r RTNode) isRoot() bool {
	// fmt.Println("isRoot", r.parent, r.parent == nil)
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

func (r RTNode) getData(i int) Rectangle {
	return r.datas[i]
}

func (r *RTNode) setDatas(index int, rect Rectangle) {
	r.datas[index] = rect
}

func (r *RTNode) addData(rect Rectangle) {
	if r.usedSpace == r.rtree.getNodeCapacity() {
		panic("Node is full.") //Node is full.
	}

	r.datas[r.usedSpace] = rect
	r.usedSpace++
}

func (r *RTNode) deleteData(i int) {
	tmp := make([]Rectangle, r.rtree.getNodeCapacity()+1)
	copy(tmp, r.datas[0:i])
	copy(tmp[i:], r.datas[i+1:])
	r.datas = tmp

	r.usedSpace--
}

// 压缩算法 叶节点L中刚刚删除了一个条目，
// 如果这个结点的条目数太少下溢， 则删除该结点，同时将该结点中剩余的条目重定位到其他结点中。
// 如果有必要，要逐级向上进行这种删除，调整向上传递的路径上的所有外包矩形，使其尽可能小，直到根节点。
// 存储删除结点中剩余条目
func (r *RTNode) condenseTree(list *[]IRTNode) {
	if r.isRoot() {
		// 根节点只有一个条目了，即只有左孩子或者右孩子 ，
		// 将唯一条目删除，释放根节点，将原根节点唯一的孩子设置为新根节点
		if !r.isLeaf() && r.usedSpace == 1 {
			child := r.getChild(0)

			child.setParent(nil)
			r.rtree.setRoot(child)
		}
	} else {
		parent := r.getParent()
		min := int(math.Round(float64(r.rtree.getNodeCapacity()) * r.rtree.fillFactor))
		if r.usedSpace < min { //递归的上一次,child进行了删除
			parent.deleteData(parent.getDeleteIndex()) // 其父节点中删除此条目
			parent.delChild(parent.getDeleteIndex())

			r.parent = nil
			*list = append(*list, r) // 之前已经把数据删除了
		} else {
			parent.setDatas(parent.getDeleteIndex(), r.getNodeRectangle())
		}
		parent.condenseTree(list)
	}
}

// --> 插入新的Rectangle后从插入的叶节点开始向上调整RTree，直到根节点
// node1 引起需要调整的孩子结点
// node2 分裂的结点，若未分裂则为null
func (r *RTNode) adjustTree(node1 IRTNode, node2 IRTNode) {
	panic("RTNode adjustTree will never be called")
}

//
// 分裂结点的平方算法
// 1、为两个组选择第一个条目--调用算法pickSeeds()来为两个组选择第一个元素，分别把选中的两个条目分配到两个组当中。<br>
// 2、检查是否已经分配完毕，如果一个组中的条目太少，为避免下溢，将剩余的所有条目全部分配到这个组中，算法终止<br>
// 3、调用pickNext来选择下一个进行分配的条目--计算把每个条目加入每个组之后面积的增量，选择两个组面积增量差最大的条目索引,
// 如果面积增量相等则选择面积较小的组，若面积也相等则选择条目数更少的组<br>
//
// @param rect
//            导致分裂的溢出Rectangle
// @return 两个组中的条目的索引
func (r *RTNode) quadraticSplit(rect Rectangle) [][]int {
	r.datas[r.usedSpace] = rect //先添加进去
	total := r.usedSpace + 1    // 结点总数

	mask := make([]int, total) // 标记访问的条目

	for i := 0; i < total; i++ {
		mask[i] = 1
	}

	minNodeSize := int(math.Round(float64(r.rtree.getNodeCapacity()) * r.rtree.fillFactor))
	// fmt.Println("minNodeSize", minNodeSize)
	if minNodeSize < 2 {
		minNodeSize = 2
	}
	// 记录没有被检查的条目的个数
	rem := total

	var group1, group2 []int

	seed := r.pickSeeds()

	group1 = append(group1, seed[0])
	mask[seed[0]] = -1

	group2 = append(group2, seed[1])
	mask[seed[1]] = -1

	rem -= 2

	for rem > 0 {
		if minNodeSize-len(group1) == rem {
			// 将剩余的所有条目全部分配到group1组中，算法终止
			for i := 0; i < total; i++ {
				if mask[i] != -1 { // 还没有被分配
					group1 = append(group1, i)
					mask[i] = -1
					rem--
				}
			}

		} else if minNodeSize-len(group2) == rem {
			// 将剩余的所有条目全部分配到group2组中，算法终止
			for i := 0; i < total; i++ {
				if mask[i] != -1 { // 还没有被分配
					group2 = append(group2, i)
					mask[i] = -1
					rem--
				}
			}

		} else {
			// 求group1中所有条目的最小外包矩形
			mbr1 := r.datas[group1[0]].clone()
			for i := 1; i < len(group1); i++ {
				mbr1 = mbr1.getUnionRectangle(r.datas[group1[i]])
			}

			// 求group2中所有条目的外包矩形
			mbr2 := r.datas[group2[0]].clone()
			for i := 1; i < len(group2); i++ {
				mbr2 = mbr2.getUnionRectangle(r.datas[group2[i]])
			}

			// 找出下一个进行分配的条目
			var dif, areaDiff1, areaDiff2 float64 = -1, -1, -1
			var sel int = -1

			for i := 0; i < total; i++ {
				if mask[i] != -1 { // 还没有被分配的条目
					// 计算把每个条目加入每个组之后面积的增量，选择两个组面积增量差最大的条目索引

					a := mbr1.getUnionRectangle(r.datas[i])
					areaDiff1 = a.getArea() - mbr1.getArea()

					b := mbr2.getUnionRectangle(r.datas[i])
					areaDiff2 = b.getArea() - mbr2.getArea()
					tmpdiff := math.Abs(areaDiff1 - areaDiff2)
					if tmpdiff > dif {
						dif = tmpdiff
						sel = i
					}
				}
			}
			if areaDiff1 < areaDiff2 { // 先比较面积增量
				group1 = append(group1, sel)
			} else if areaDiff1 > areaDiff2 {
				group2 = append(group2, sel)
			} else if mbr1.getArea() < mbr2.getArea() { // 再比较自身面积
				group1 = append(group1, sel)
			} else if mbr1.getArea() > mbr2.getArea() {
				group2 = append(group2, sel)
			} else if len(group1) < len(group2) { // 最后比较条目个数
				group1 = append(group1, sel)
			} else if len(group1) > len(group2) {
				group2 = append(group2, sel)
			} else { //随便
				group1 = append(group1, sel)
			}
			mask[sel] = -1
			rem--

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
	var i1, i2 int = 0, 1

	for i := 0; i < r.usedSpace; i++ {
		for j := i + 1; j <= r.usedSpace; j++ {
			rect := r.datas[i].getUnionRectangle(r.datas[j])
			d := rect.getArea() - r.datas[i].getArea() - r.datas[j].getArea()

			if i == 0 && j == 1 {
				inefficiency = d
				continue
			}

			if d > inefficiency {
				inefficiency = d
				i1 = i
				i2 = j
			}
		}
	}
	if i1 == i2 {
		panic("pick seed err")
	}
	return []int{i1, i2}
}

func (r RTNode) getNodeRectangle() Rectangle {
	if r.usedSpace > 0 {
		return getUnionRectangle(r.datas, r.usedSpace)
	} else if r.usedSpace == 0 {
		panic("empty RTNode")
	}
	return Rectangle{}
}

// 步骤CL1：初始化――记R树的根节点为N。
// 步骤CL2：检查叶节点――如果N是个叶节点，返回N
// 步骤CL3：选择子树――如果N不是叶节点，则从N中所有的条目中选出一个最佳的条目F
// 选择的标准是：如果E加入F后，F的外廓矩形FI扩张最小，则F就是最佳的条目。如果有两个
// 条目在加入E后外廓矩形的扩张程度相等，则在这两者中选择外廓矩形较小的那个。
// 步骤CL4：向下寻找直至达到叶节点――记Fp指向的孩子节点为N，然后返回步骤CL2循环运算， 直至查找到叶节点。
//
// @param Rectangle
// @return RTDataNode
func (r RTNode) chooseLeaf(rect Rectangle) *RTDataNode {
	defer panic("RTNode chooseLeaf will never be called")
	return &RTDataNode{}
}

// R树的根节点为T，查找包含rectangle的叶子结点
//
// 1、如果T不是叶子结点，则逐个查找T中的每个条目是否包围rectangle，若包围则递归调用findLeaf()
// 2、如果T是一个叶子结点，则逐个检查T中的每个条目能否匹配rectangle
//
// @param rectangle
// @return 返回包含rectangle的叶节点
func (r RTNode) findLeaf(rect Rectangle) *RTDataNode {
	defer panic("RTNode findLeaf will never be called")
	return &RTDataNode{}
}
