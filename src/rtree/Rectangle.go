package rtree

// Rectangle 外接矩形
type Rectangle struct {
	low  Point
	high Point
}

func (r Rectangle) clone() (c Rectangle) {
	c.low = r.low.clone()
	c.high = r.high.clone()
	return
}

func (r *Rectangle) init(low *Point, high *Point) {
	if low.getDimension() != high.getDimension() {
		return
	}

	for i, d := range low.data {
		if d > high.getFloatCoordinate(i) {
			return
		}
	}

	r.low = low.clone()
	r.high = high.clone()
}

func (r Rectangle) getLow() Point {
	return r.low.clone()
}

func (r Rectangle) getHigh() Point {
	return r.high.clone()
}

func (r Rectangle) getUnionRectangle(in Rectangle) (out Rectangle) {

	//比较一个纬度即可
	if r.low.getDimension() != in.low.getDimension() {
		return
	}

	low := r.low.clone()
	high := r.high.clone()

	for i, d := range in.low.data {

		if d < low.getFloatCoordinate(i) {
			low.data[i] = d
		}

		if in.high.data[i] > high.getFloatCoordinate(i) {
			high.data[i] = in.high.getFloatCoordinate(i)
		}
	}

	out.low = low
	out.high = high

	return
}

func (r Rectangle) getArea() (area float64) {
	area = 1

	for i, d := range r.low.data {
		area *= r.high.getFloatCoordinate(i) - d
	}

	return area
}

func getUnionRectangleForMany(ins []Rectangle) (out Rectangle) {
	if len(ins) == 0 {
		return
	}

	for i, r := range ins {
		if i == 0 {
			out = r.clone()
		}

	}

}
