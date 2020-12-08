package rtree

// RTree ...
type RTree struct {
	root         *RTNode // 根节点
	treeType     int     // 树类型
	nodeCapacity int     // 结点容量
	fillFactor   float64 // 结点填充因子 ，用于计算每个结点最小条目个数
	dimension    int     // 维度
}

func (r *RTree) init(capacity int, fillFactor float64, t int, dimension int) {
	r.fillFactor = fillFactor
	r.nodeCapacity = capacity
	r.dimension = dimension
	r.root = new(RTNode)
}

func (r RTree) getDimension() int {
	return r.dimension
}

func (r *RTree) insert(rect Rectangle) {

}

func (r RTree) getNodeCapacity() int {
	return r.nodeCapacity
}
