package rtree

import (
	"fmt"
)

// Point n维空间中的点，所有的维度被存储在一个float数组中
type Point struct {
	data []float64
}

func (p Point) toString() (ret string) {
	for i, d := range p.data {
		if i != 0 {
			ret += ","
		}
		ret += fmt.Sprintf("%f", d)
	}

	return
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
	}
	return 0
}

func (p Point) equals(o Point) bool {
	if p.getDimension() != o.getDimension() {
		return false
	}

	for i, d := range p.data {
		if o.data[i] != d {
			return false
		}
	}

	return true
}
