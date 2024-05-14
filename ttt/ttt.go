package ttt

import (
	"bufio"
	"fmt"
	"os"
)

func Play(options *Options) {
	reader := bufio.NewReader(os.Stdout)
	grid := init_grid(options.Token1, options.Token2, options.Empty, options.Dimension)

	// initial screen
	print_grid(grid)

	player := "o"
	var s1, s2 string
	for {
		// take user input and clear screen
		fmt.Printf("\n\n\n\n\n\n\t\t-<< %s >>-\n\n\tinput>_ ", player)
		s, _ := reader.ReadString(10)
		ClearScreen()

		// process input
		_, err := fmt.Sscan(s, &s1, &s2)
		if err != nil {
			fmt.Printf("[Error] Invalid input '%s %s'\n", s1, s2)
			return
		}
		if !(len(s1) == len(s2) && len(s1) == 2) {
			fmt.Printf("[Error] Invalid input format")
			return
		}
		// subtract 1 to account for the fact that our grid on screen is 1-indexed
		point1 := Point{int(s1[0] - '1'), int(s1[1] - '1')}
		point2 := Point{int(s2[0] - '1'), int(s2[1] - '1')}
		fmt.Println(point1, point2)

		print_grid(grid)

		if player == "o" {
			player = "x"
			update_grid(grid, point1, point2, options.Empty, options.Token1)
		} else {
			player = "o"
			update_grid(grid, point1, point2, options.Empty, options.Token2)
		}

	}
}
