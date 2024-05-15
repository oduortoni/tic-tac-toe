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
			if row != 0 && grid[row][0] == options.TokTop {
				return true, options.TokTop
			} else if row != dimension-1 && grid[row][0] == options.TokBtm {
				return true, options.TokBtm
			}
		}

		// complete column
		if countHoriz == dimension-1 {
			if grid[0][col] == options.TokTop {
				return true, options.TokTop
			} else {
				return true, options.TokBtm
			}
		}

		row++
		col++
		countVert = 0
		countHoriz = 0
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

	// can only move one square a time
	if (dx == 0 || dx == 1) && (dy == 1 || dy == 0) {
		valid = true
	} else {
		valid = false
	}

	return valid
}
