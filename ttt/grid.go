package ttt

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

func print_grid(grid Grid, options *Options) {
	size := len(grid)
	color := ORIGINAL
	boxNumber := 1

	print("\n\n")
	for i := 0; i < size; i++ {
		for k := 0; k < 3; k++ {
			print("\t")
			for j := 0; j < size; j++ {
				for l := 0; l < 5; l++ {
					token := grid[i][j]
					if token == options.SymbolTop {
						color = options.ColorTop
					} else if token == options.SymbolBtm {
						color = options.ColorBtm
					} else {
						color = ORIGINAL
					}
					if k == 1 && l == 1 {
						print(MAGENTA)
						print(boxNumber)
						print(ORIGINAL)
						boxNumber++
					} else if k == 2 && l == 2 {
						word := grid[i][j]
						print(color + word + ORIGINAL)
					} else {
						print(" ")
					}
				}
			}
			print("\n")
		}
	}
	// println()
}
