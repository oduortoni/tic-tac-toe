package ttt

import "fmt"

func init_grid(tok1, tok2, empty string, dim int) (grid Grid) {
	grid = make(Grid, dim)

	tok := empty
	for i := 0; i < dim; i++ {
		if i == 0 {
			tok = tok1
		} else if i == dim-1 {
			tok = tok2
		} else {
			tok = empty
		}
		for j := 0; j < dim; j++ {
			grid[i] = append(grid[i], tok)
		}
	}
	return
}

func print_grid(grid Grid, options *Options) {
	size := len(grid)

	fmt.Printf("\n\n\n\n")
	line := 1
	
	// top column numbers
	print("\t\t+ ")
	for i := 0; i < size; i++ {
		print(" ", i+1, " ")
	}
	print(" + ")
	println()
	println()

	color := ORIGINAL
	for i := 0; i < size; i++ {
		// left line numbers
		print("\t\t", line, " ")
		line++
		for j := 0; j < size; j++ {
			token := grid[i][j]
			if token == options.Token1 {
				color = options.Color1
			} else if token == options.Token2 {
				color = options.Color2
			} else {
				color = ORIGINAL
			}
			word := " " + grid[i][j] + " "
			print(color + word + ORIGINAL)
		}
		// right line numbers
		print(" ", line)
		print("\n")
		println()
	}

	// bottom column numbers
	print("\t\t+ ")
	for i := 0; i < size; i++ {
		print(" ", i+1, " ")
	}
	print(" + ")
	println()
	
}

func update_grid(grid Grid, p1, p2 Point, empty, token string) (updated bool) {
	width := len(grid[0])
	height := len(grid)

	if validMove := check_valid(p1, p2, width, height); validMove {
		// update iff the square to move to is empty
		if grid[p2.X][p2.Y] == empty {
			grid[p1.X][p1.Y] = empty
			grid[p2.X][p2.Y] = token
			updated = true
		}
	}

	return
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
