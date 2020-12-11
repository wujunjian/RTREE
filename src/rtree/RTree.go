package rtree

// RTree ...
type RTree struct {
	root         IRTNode // 根节点
	treeType     int     // 树类型
	nodeCapacity int     // 结点容量
	fillFactor   float64 // 结点填充因子 ，用于计算每个结点最小条目个数
	dimension    int     // 维度
}

func (r *RTree) Init(capacity int, fillFactor float64, t int, dimension int) {
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

func (r *RTree) Insert(rect Rectangle) bool {
	if rect.isNULL() {
		panic("Rectangle cannot be null.")
	}
	if rect.getDimension() != r.getDimension() {
		panic("Rectangle dimension different than RTree dimension.")
	}

	leaf := r.root.chooseLeaf(rect)

	return leaf.insert(rect)
}

type LngLatPoint struct {
	Lng float64
	Lat float64
}

func CoordinatesToRectangle(points []LngLatPoint, info interface{}) Rectangle {
	var minlng, minlat, maxlng, maxlat float64

	for i, p := range points {
		if i == 0 {
			minlng = p.Lng
			maxlng = p.Lng
			minlat = p.Lat
			maxlat = p.Lat
		}
		if p.Lng < minlng {
			minlng = p.Lng
		}
		if p.Lng > maxlng {
			maxlng = p.Lng
		}

		if p.Lat < minlat {
			minlat = p.Lat
		}
		if p.Lat > maxlat {
			maxlat = p.Lat
		}
	}

	var rect Rectangle
	rect.low.data = append(rect.low.data, minlng, minlat)
	rect.high.data = append(rect.high.data, maxlng, maxlat)
	rect.info = info

	return rect
}

func (r *RTree) InsertCoordinates(points []LngLatPoint, info interface{}) bool {
	return r.Insert(CoordinatesToRectangle(points, info))
}

func (r *RTree) SearchCoordinates(points []LngLatPoint) []Rectangle {
	return r.Search(CoordinatesToRectangle(points, nil))
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
func (r *RTree) Delete(rect Rectangle) int {
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

func (r RTree) Search(rect Rectangle) []Rectangle {
	return r.root.Search(rect)
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

func (r RTree) Range() (leaf []Rectangle) {

	node := r.root
	leaf = r.RangeDir(node)
	return
}

func (r RTree) RangeDir(node IRTNode) (leaf []Rectangle) {
	if node.isLeaf() {
		for i := 0; i < node.getUsedSpace(); i++ {
			leaf = append(leaf, node.getData(i))
		}
	} else {
		for i := 0; i < node.getUsedSpace(); i++ {
			leaf = append(leaf, r.RangeDir(node.getChild(i))...)
		}
	}

	return
}
