package ttt

func toCoord(cart map[int]*Point, key int) (bool, *Point) {
	point, present := cart[key]
	return present, point
}

func toCoordinateSystem(dimension int) map[int]*Point {
	coord := make(map[int]*Point, dimension)
	num := 1
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			coord[num] = &Point{i, j}
			num++
		}
	}
	return coord
}
