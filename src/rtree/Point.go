package rtree

// Point n维空间中的点，所有的维度被存储在一个float数组中
type Point struct {
	data []float64
}

func (p Point) toString() {

}

func (p Point) getDimension() int {
	return len(p.data)
}
