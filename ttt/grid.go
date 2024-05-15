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
			if token == options.TokTop {
				color = options.ColorTop
			} else if token == options.TokBtm {
				color = options.ColorBtm
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

/*
*
* before motion, check whether a move is valid
* update iff the square to move to is empty
* ensure player can only move their own token
*
*/
func update_grid(grid Grid, p1, p2 Point, empty, token string) (updated bool) {
	width := len(grid[0])
	height := len(grid)

	if validMove := check_valid(p1, p2, width, height); validMove {
		if grid[p2.X][p2.Y] == empty && grid[p1.X][p1.Y] == token {
			grid[p1.X][p1.Y] = empty
			grid[p2.X][p2.Y] = token
			updated = true
		}
	}

	return
}
