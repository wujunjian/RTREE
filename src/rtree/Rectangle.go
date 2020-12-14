package rtree

// Rectangle 外接矩形
type Rectangle struct {
	low  Point
	high Point
	info interface{}
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
			panic(1)
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

func (r Rectangle) toString() string {
	return r.low.toString() + ";" + r.high.toString()
}

func (r Rectangle) getUnionRectangle(in Rectangle) (out Rectangle) {

	//比较一个纬度即可
	if r.low.getDimension() != in.low.getDimension() {
		panic("getUnionRectangle error Dimension")
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
		area *= (r.high.getFloatCoordinate(i) - d)
	}

	return area
}

// 两个Rectangle相交的面积
func (r Rectangle) intersectingArea(o Rectangle) (ret float64) {
	if !r.isIntersection(o) { //不相交
		return 0
	}
	ret = 1
	for i, lowd := range r.low.data {
		l1 := lowd
		h1 := r.high.getFloatCoordinate(i)
		l2 := o.low.getFloatCoordinate(i)
		h2 := o.high.getFloatCoordinate(i)
		if l1 <= l2 && h1 <= h2 {
			ret *= (h1 - l2)
		} else if l1 >= l2 && h1 >= h2 {
			ret *= (h2 - l1)
		} else if l1 >= l2 && h1 <= h2 {
			ret *= (h1 - l1)
		} else if l1 <= l2 && h1 >= h2 {
			ret *= (h2 - l2)
		}
	}

	return
}

// 判断两个Rectangle是否相交
func (r Rectangle) isIntersection(o Rectangle) bool {

	for i, lowd := range r.low.data {
		if lowd > o.high.getFloatCoordinate(i) ||
			r.high.getFloatCoordinate(i) < o.low.getFloatCoordinate(i) {
			return false
		}
	}
	return true
}

// 判断 入参:rect 是否被包围
func (r Rectangle) enclosure(o Rectangle) bool {
	if o.isNULL() {
		panic("Rectangle cannot be null.")
	}

	if r.getDimension() != o.getDimension() {
		panic("Rectangle dimension is different from current dimension.")
	}

	// 只要传入的rectangle有一个维度的坐标越界了就不被包含
	for i := 0; i < r.getDimension(); i++ {

		if o.low.getFloatCoordinate(i) < r.low.getFloatCoordinate(i) ||
			o.high.getFloatCoordinate(i) > r.high.getFloatCoordinate(i) {
			return false
		}
	}

	return true
}

func (r Rectangle) getDimension() int {
	return len(r.low.data)
}

func (r Rectangle) isNULL() bool {
	if len(r.low.data) == 0 {
		return true
	}
	return false
}

func (r *Rectangle) clean() {
	r.low.data = nil
	r.high.data = nil
}

func (r Rectangle) equals(o Rectangle) bool {
	if r.low.equals(o.low) && r.high.equals(o.high) {
		return true
	}
	return false
}

func getUnionRectangle(ins []Rectangle, usedSpace int) (out Rectangle) {
	if len(ins) == 0 {
		return
	}

	for i, r := range ins {
		if i >= usedSpace {
			break
		}

		if i == 0 {
			out = r.clone()
			continue
		}

		out = out.getUnionRectangle(r)
	}
	return
}
