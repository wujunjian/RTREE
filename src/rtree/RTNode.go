package rtree

// RTNode ...
type RTNode struct {
	rtree *RTree // 结点所在的树
	level int // 结点所在的层
	datas []Rectangle // 相当于条目
	parent *RTNode // 父节点
	usedSpace int // 结点已用的空间
	insertIndex int // 记录插入的搜索路径索引
	deleteIndex int // 记录删除的查找路径索引
}
