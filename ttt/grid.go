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

func print_grid(grid Grid) {
	size := len(grid)

	fmt.Printf("\n\n\n\n")
	line := 1
	// column numbers
	print("\t\t  ")
	for i := 0; i < size; i++ {
		print(" ", i+1, " ")
	}
	println()
	println()

	for i := 0; i < size; i++ {
		print("\t\t", line, " ")
		line++
		for j := 0; j < size; j++ {
			print(" ", grid[i][j], " ")
		}
		print("\n")
		println()
	}
}

func update_grid(grid Grid, p1, p2 Point, empty, token string) (updated bool) {
	grid[p1.Y][p1.X] = empty
	grid[p2.Y][p2.X] = token

	updated = true
	return
}
