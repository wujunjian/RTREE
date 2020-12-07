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

func (p Point) clone() Point {

	var r Point

	for _, d := range p.data {
		r.data = append(r.data, d)
	}

	return r
}

func (p Point) getFloatCoordinate(index int) float64 {
	if p.getDimension() > index {
		return p.data[index]
	} else {
		return 0
	}
}