package rtree

// RTree ...
type RTree struct {
	root         IRTNode // 根节点
	treeType     int     // 树类型
	nodeCapacity int     // 结点容量
	fillFactor   float64 // 结点填充因子 ，用于计算每个结点最小条目个数
	dimension    int     // 维度
}

func (r *RTree) init(capacity int, fillFactor float64, t int, dimension int) {
	r.fillFactor = fillFactor
	r.nodeCapacity = capacity
	r.treeType = t
	r.dimension = dimension

	dataNode := new(RTDataNode)
	dataNode.init(r, nil)
	r.root = dataNode
}

func (r RTree) getDimension() int {
	return r.dimension
}

func (r *RTree) insert(rect Rectangle) bool {
	if rect.isNULL() {
		panic(1)
	}
	if rect.getDimension() != r.getDimension() {
		panic(1)
	}

	leaf := r.root.chooseLeaf(rect)

	return leaf.insert(rect)
}

func (r RTree) getNodeCapacity() int {
	return r.nodeCapacity
}

func (r *RTree) setRoot(root IRTNode) {
	r.root = root
}

//  从R树中删除Rectangle
//  1、寻找包含记录的结点--调用算法findLeaf()来定位包含此记录的叶子结点L，如果没有找到则算法终止。
//  2、删除记录--将找到的叶子结点L中的此记录删除
//  3、调用算法condenseTree
//  @param rectangle
//  @return
func (r *RTree) delete(rect Rectangle) int {
	if rect.isNULL() {
		panic("Rectangle cannot be null.")
	}

	if rect.getDimension() != r.getDimension() {
		panic("Rectangle dimension different than RTree dimension.")
	}

	if leaf := r.root.findLeaf(rect); leaf != nil {
		return leaf.delete(rect)
	}
	return -1
}

func traversePostOrder(root IRTNode) (list []IRTNode) {
	if root == nil {
		panic("Node cannot be null.")
	}
	list = append(list, root)

	if !root.isLeaf() {
		for i := 0; i < root.dataLength(); i++ {
			list = append(list, traversePostOrder(root.getChild(i))...)
		}
	}

	return
}
