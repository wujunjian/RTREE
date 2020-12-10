package rtree

// RTDataNode 叶子结点
type RTDataNode struct {
	RTNode
}

func (r *RTDataNode) init(rtree *RTree, paraent IRTNode) {
	r.RTNode.init(rtree, paraent, 0)
}

func (r *RTDataNode) insert(rect Rectangle) bool {
	if r.usedSpace < r.rtree.getNodeCapacity() { // 已用节点小于节点容量
		r.datas[r.usedSpace] = rect
		r.usedSpace++

		if r.parent != nil { //调整树，但不需要分裂节点，因为 节点小于节点容量，还有空间
			r.parent.adjustTree(r, nil)
		}
		return true
	} else { // 超过结点容量

	}

	return true
}

func (r *RTDataNode) splitLeaf(rect Rectangle) []RTDataNode {
	var group [][]int

	switch r.rtree.treeType {
	case RTREE_LINEAR:
	case RTREE_QUADRATIC:
		group = r.quadraticSplit(rect)
	case RTREE_EXPONENTIAL:
	case RSTAR:
	default:
		panic(1)
	}

	var l, ll RTDataNode
	l.init(r.rtree, r.parent)
	ll.init(r.rtree, r.parent)

	for i, g := range group {
		if i == 0 {
			for _, d := range g {
				l.addData(r.datas[d])
			}
		} else {
			for _, d := range g {
				ll.addData(r.datas[d])
			}
		}
	}

	return []RTDataNode{l, ll}
}

//@override
func (r *RTDataNode) chooseLeaf(rect Rectangle) *RTDataNode {
	r.insertIndex = r.usedSpace

	return r
}
