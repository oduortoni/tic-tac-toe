package ttt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Play(options *Options) {
	// begin with an empty screen
	ClearScreen()

	reader := bufio.NewReader(os.Stdout)
	grid := init_grid(options.TokTop, options.TokBtm, options.Empty, options.Dimension)

	// initial screen
	print_grid(grid, options)

	player := "o"
	var s1, s2 string
	for {
		// take user input and clear screen
		fmt.Printf("\n\n\n\n\n\n\t\t-<< %s >>-\n\n\tinput>_ ", player)
		s, _ := reader.ReadString(10)
		s = strings.TrimSpace(s)

		// quit game
		if s == "quit" || s == "q" || s == "exit" || s == "e" {
			print_success(grid, player)

			//print_exit(grid, player)
			print_grid(grid, options)
			fmt.Printf("\n\n\n\n\n\n")
			return
		}

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
		//fmt.Println(point1, point2)

		print_grid(grid, options)

		if player == options.TokTop {
			if updated := update_grid(grid, point1, point2, options.Empty, options.TokTop); updated {
				player = options.TokBtm
			}
		} else {
			if updated := update_grid(grid, point1, point2, options.Empty, options.TokBtm); updated {
				player = options.TokTop
			}
		}
		// check for winner
		if found, who := winner(grid, options); found {
			print_success(grid, who)
			print_grid(grid, options)
			print("\n\n\n\n\n")
			return
		}

		print_grid(grid, options)
	}
}
