package rtree

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

func (r *RTNode) addData(rect Rectangle) {
	if r.usedSpace == r.rtree.getNodeCapacity() {
		return //Node is full.
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
