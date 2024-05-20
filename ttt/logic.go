package ttt

func winner(grid Grid, options *Options) (bool, string) {
	dimension := len(grid)

	countVert := 0
	countHoriz := 0
	row := 0
	col := 0
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension-1; j++ {
			if grid[row][j] == grid[row][j+1] {
				countVert++
			}
			if grid[j][col] == grid[j+1][col] {
				countHoriz++
			}
		}
		println()

		// complete row bar starting position
		if countVert == dimension-1 {
			if row != 0 && grid[row][0] == options.SymbolTop {
				return true, options.SymbolTop
			} else if row != dimension-1 && grid[row][0] == options.SymbolBtm {
				return true, options.SymbolBtm
			}
		}

		// complete row
		if countHoriz == dimension-1 {
			if grid[0][col] == options.SymbolBtm {
				return true, options.SymbolBtm
			} else {
				return true, options.SymbolTop
			}
		}

		row++
		col++
		countVert = 0
		countHoriz = 0
	}

	// check for a winner within the diagonal
	// leading diagonal
	prev := grid[0][0]
	all := true
	for i := 1; i < dimension; i++ {
		for j := 1; j < dimension; j++ {
			if i == j && !(grid[i][j] == prev) {
				all = false
				break
			}
		}
	}

	if all && prev != options.Empty{
		return true, prev
	}

	// following diagonal
	prev = grid[0][dimension-1]
	all = true
	for i := dimension-1; i >= 0; i-- {
		for j := 0; j < dimension; j++ {
			if (i+j) == dimension-1 && !(grid[i][j] == prev) {
				all = false
				break
			}
		}
	}

	if all && prev != options.Empty {
		return true, prev
	}

	return false, ""
}

func check_valid(p1, p2 Point, width, height int) bool {
	valid := true
	// check bounds
	if (p1.X < 0 || p1.X >= width || p2.X < 0 || p2.X >= width) || (p1.Y < 0 || p1.Y >= height || p2.Y < 0 || p2.Y >= height) { // outside
		return false
	}

	if p2.X > p1.X && p2.X > width { // l to r
		valid = false
	}
	if p2.X < p1.X && p2.X < 0 { // r to l
		valid = false
	}
	if p2.Y > p1.Y && p2.Y > height { // t to b
		valid = false
	}
	if p2.Y < p1.Y && p2.Y < 0 { // b to t
		valid = false
	}

	// distance to move
	dx := p2.X - p1.X
	if dx < 0 {
		dx *= -1
	}
	dy := p2.Y - p1.Y
	if dy < 0 {
		dy *= -1
	}

	// can only move one square a time (r l t b)
	delta := dy + dx
	if delta == 1 {
		valid = true
	} else {
		valid = false
	}

	return valid
}

func computeCoords(grid Grid, space, token string) (*Point, *Point) {
	dimension := len(grid)

	point1 := Point{0, 0}
	point2 := Point{0, 0}

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if grid[i][j] == token {
				if i < dimension-1 && grid[i+1][j] == space {
					point1.X = i
					point1.Y = j
					point2.X = i + 1
					point2.Y = j
					break
				}
				if j < dimension-1 && grid[i][j+1] == space {
					point1.X = i
					point1.Y = j
					point2.X = i
					point2.Y = j + 1
					break
				}
				if i > 0 && grid[i-1][j] == space {
					point1.X = i
					point1.Y = j
					point2.X = i - 1
					point2.Y = j
					break
				}
				if j > 0 && grid[i][j-1] == space {
					point1.X = i
					point1.Y = j
					point2.X = i
					point2.Y = j - 1
					break
				}
			}
		}
	}
	return &point1, &point2
}
