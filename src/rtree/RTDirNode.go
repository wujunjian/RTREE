package rtree

import (
	"math"
)

// RTDirNode 非叶节点
type RTDirNode struct {
	RTNode
	children []IRTNode //孩子结点集
}

func (r *RTDirNode) init(rtree *RTree, parent IRTNode, level int) {
	//	r.children
	r.RTNode.init(rtree, parent, level)
}

func (r RTDirNode) getChild(index int) IRTNode {

	return r.children[index]
}

//@override
func (r RTDirNode) chooseLeaf(rect Rectangle) *RTDataNode {
	var index int

	switch r.rtree.treeType {
	case RTREE_LINEAR:
		fallthrough
	case RTREE_QUADRATIC:
		fallthrough
	case RTREE_EXPONENTIAL:
		index = r.findLeastEnlargement(rect)
	case RSTAR:
		if r.level == 1 {
			index = r.findLeastOverlap(rect) // 获得最小重叠面积的结点的索引
		} else {
			index = r.findLeastEnlargement(rect) // 获得面积增量最小的结点的索引
		}

	default:
		panic(1)
	}
	r.insertIndex = index

	return r.getChild(index).chooseLeaf(rect) // 非叶子节点的chooseLeaf（）实现递归调用

	return &RTDataNode{}
}

// @param rectangle
// @return -->返回最小重叠面积的结点的索引， 如果重叠面积相等则选择加入此Rectangle后面积增量更小的，
//         如果面积增量还相等则选择自身面积更小的
func (r RTDirNode) findLeastOverlap(rect Rectangle) int {
	var overlap float64 = -1
	var sel int = -1

	for i := 0; i < r.usedSpace; i++ {
		node := r.getChild(i)
		var ol float64 // 用于记录每个孩子的datas数据与传入矩形的重叠面积之和

		for j := 0; j < len(node.datas); j++ {
			// 将传入矩形与各个矩形重叠的面积累加到ol中，得到重叠的总面积
			ol += rect.intersectingArea(node.datas[j])
		}

		if ol < overlap {
			overlap = ol // 记录重叠面积最小的
			sel = i      //记录第几个孩子的索引
		} else if ol == overlap {
			// 如果重叠面积相等则选择加入此Rectangle后面积增量更小的,
			// 如果面积增量还相等则选择自身面积更小的

			area1 := r.datas[i].getUnionRectangle(rect).getArea() - r.datas[i].getArea()
			area2 := r.datas[sel].getUnionRectangle(rect).getArea() - r.datas[sel].getArea()

			if area1 == area2 {
				if r.datas[sel].getArea() > r.datas[i].getArea() {
					sel = i
				}
			} else if area1 < area2 {
				sel = i
			}
		}

	}
	return sel
}

// @param rectangle
//  @return -->面积增量最小的结点的索引，如果面积增量相等则选择自身面积更小的
func (r RTDirNode) findLeastEnlargement(rect Rectangle) int {
	var area float64 = math.MaxFloat64
	var sel int = -1

	for i := 0; i < r.usedSpace; i++ {
		// 增量enlargement = 包含（datas[i]里面存储的矩形与查找的矩形）的最小矩形的面积
		// datas[i]里面存储的矩形的面积
		enlargement := r.datas[i].getUnionRectangle(rect).getArea() - r.datas[i].getArea()
		if enlargement < area {
			area = enlargement // 记录增量
			sel = i
		} else if enlargement == area {
			if r.datas[sel].getArea() >= r.datas[i].getArea() {
				sel = i
			}
		}
	}
	return sel
}

//  --> 插入新的Rectangle后从插入的叶节点开始向上调整RTree，直到根节点
//  @param node1
//             引起需要调整的孩子结点
//  @param node2
//             分裂的结点，若未分裂则为null
func (r *RTDirNode) adjustTree(node1, node2 IRTNode) {
	// 先要找到指向原来旧的结点（即未添加Rectangle之前）的条目的索引
	r.datas[r.insertIndex] = node1.getNodeRectangle() // 先用node1覆盖原来的结点
	r.children[r.insertIndex] = node1                 // 替换旧的结点

	if node2 != nil {
		r.insert(node2) // 插入新的结点
	} else if !r.isRoot() {
		r.parent.adjustTree(r, nil)
	}
}

//  -->非叶子节点插入
//  @param node
//  @return 如果结点需要分裂则返回true
func (r *RTDirNode) insert(node IRTNode) bool {
	if r.usedSpace < r.rtree.getNodeCapacity() {

		r.datas[r.usedSpace] = node.getNodeRectangle()
		r.usedSpace++

		r.children = append(r.children, node) // new
		node.setParent(r)                     // new

		if r.parent != nil { // 不是根节点
			r.parent.adjustTree(r, nil)
		}
		return false
	} else { // 非叶子结点需要分裂
		a := r.splitIndex(node)
	}

	return true
}

func (r *RTDirNode) splitIndex(node IRTNode) []IRTNode {

	var group [][]int

	switch r.rtree.treeType {
	case RTREE_LINEAR:
	case RTREE_QUADRATIC:
	case RTREE_EXPONENTIAL:
		group = r.quadraticSplit(node.getNodeRectangle())
		r.children = append(r.children, node) // new
		node.setParent(r)                     // new
	case RTREE_EXPONENTIAL:
	case RSTAR:
	default:
		panic(1)
	}
	var index1, index2 RTDirNode
	index1.init(r.rtree, r.parent, r.level)
	index2.init(r.rtree, r.parent, r.level)

	group1 := group[0]
	group2 := group[1]

	for _, i := range group1 {
		index1.datas = append(index1.datas, r.datas[i])
		index1.children = append(index1.children, r.children[i])

		r.children[i].setParent(&index1)
	}

	for _, i := range group2 {
		index2.datas = append(index2.datas, r.datas[i])
		index2.children = append(index2.children, r.children[i])

		r.children[i].setParent(&index2)
	}

	return []IRTNode{&index1, &index2}
}
