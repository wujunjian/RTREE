package rtree

// RTDataNode 叶子结点
type RTDataNode struct {
	RTNode
}

func (r *RTDataNode) init(rtree *RTree, paraent IRTNode) {
	r.RTNode.init(rtree, paraent, 0)
}

//  -->叶节点中插入Rectangle 在叶节点中插入Rectangle，插入后如果其父节点不为空则需要向上调整树直到根节点；
//  如果其父节点为空，则是从根节点插入 若插入Rectangle之后超过结点容量则需要分裂结点 【注】插入数据后，从parent处开始调整数据
//
//  @param rectangle
//  @return
func (r *RTDataNode) insert(rect Rectangle) bool {
	if r.usedSpace < r.rtree.getNodeCapacity() { // 已用节点小于节点容量
		r.datas[r.usedSpace] = rect
		r.usedSpace++

		if r.parent != nil { //调整树，但不需要分裂节点，因为 节点小于节点容量，还有空间
			r.parent.adjustTree(r, nil)
		}
		return true
	} else { // 超过结点容量
		splitNodes := r.splitLeaf(rect)
		l := splitNodes[0]
		ll := splitNodes[1]

		if r.isRoot() {
			// 根节点已满，需要分裂。创建新的根节点
			var dirNode RTDirNode
			dirNode.init(r.rtree, nil, r.level+1)

			r.rtree.setRoot(&dirNode)
			// getNodeRectangle()返回包含结点中所有条目的最小Rectangle
			dirNode.addData(l.getNodeRectangle())
			dirNode.addData(ll.getNodeRectangle())

			l.setParent(&dirNode)
			ll.setParent(&dirNode)

			dirNode.children = append(dirNode.children, &l, &ll)
		} else { // 不是根节点
			r.parent.adjustTree(&l, &ll)
		}
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

// 从叶节点中删除此条目rectangle
// 先删除此rectangle，再调用condenseTree()返回删除结点的集合，把其中的叶子结点中的每个条目重新插入；
// 非叶子结点就从此结点开始遍历所有结点，然后把所有的叶子结点中的所有条目全部重新插入
// @param rectangle
// @return
func (r *RTDataNode) delete(rect Rectangle) int {
	for i := 0; i < r.usedSpace; i++ {
		if r.datas[i].equals(rect) {
			r.deleteData(i)
			var deleteEntriesList []IRTNode
			r.condenseTree(deleteEntriesList)

			// 重新插入删除结点中剩余的条目
			for _, node := range deleteEntriesList {
				if node.isLeaf() { // 叶子结点，直接把其上的数据重新插入
					for k := 0; k < node.dataLength(); k++ {
						r.rtree.insert(node.getData(k))
					}
				} else { // 非叶子结点，需要先后序遍历出其上的所有结点
					traverseNodes := traversePostOrder(node)

					for _, traverseNode := range traverseNodes {
						if traverseNode.isLeaf() {
							for t := 0; t < traverseNode.dataLength(); t++ {
								r.rtree.insert(traverseNode.getData(t))
							}
						}
					}
				}
			}
			return r.deleteIndex
		} // end if
	} // end for
	return -1
}

// @Override
func (r *RTDataNode) findLeaf(rect Rectangle) *RTDataNode {

	for i, d := range r.datas {
		if d.enclosure(rect) {
			r.deleteIndex = i
			return r
		}
	}
	return nil
}

//@override
func (r RTDataNode) Search(rect Rectangle, leaf []Rectangle) {
	for _, d := range r.datas {
		if rect.enclosure(d) { //d 被包含 或完全相等
			leaf = append(leaf, d)
		} else if d.enclosure(rect) { //rect 被包含,查找矩形小于条目
			leaf = append(leaf, d)
		} else if d.isIntersection(rect) { //有部分交集? 是否返回
			leaf = append(leaf, d)
		}
	}
}
