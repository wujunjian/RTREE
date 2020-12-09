package rtree

// RTDirNode 非叶节点
type RTDirNode struct {
	RTNode
	children []RTNode //孩子结点集
}

func (r *RTDirNode) init(rtree *RTree, parent *RTNode, level int) {
	//	r.children
	r.RTNode.init(rtree, parent, level)
}

func (r RTDirNode) getChild(index int) *RTNode {

	return &r.children[index]
}
