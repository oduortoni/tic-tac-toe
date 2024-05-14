package main

func main() {
	dim := 3
	token1 := "X"
	token2 := "O"
	empty := "."

	grid := init_grid(token1, token2, empty, dim)
	print_grid(grid)
}
func init_grid(tok1, tok2, empty string, dim int) (grid [][]string) {
	grid = make([][]string, dim)

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
func print_grid(grid [][]string) {
	size := len(grid)

	line:= 1
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
