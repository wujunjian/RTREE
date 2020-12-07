package rtree

// RTree ...
type RTree struct {
	root         *RTNode // 根节点
	treeType     int     // 树类型
	nodeCapacity int     // 结点容量
	fillFactor   float64 //  结点填充因子 ，用于计算每个结点最小条目个数
	dimension    int     // 维度
}
