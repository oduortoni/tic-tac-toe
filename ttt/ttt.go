package ttt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Play(options *Options) {
	coordinates := toCoordinateSystem(options.Dimension)

	// begin with an empty screen
	ClearScreen()

	reader := bufio.NewReader(os.Stdout)
	grid := init_grid(options.TokTop, options.TokBtm, options.Empty, options.Dimension)

	// initial screen
	print_grid(grid, options)

	player := options.TokTop
	for {
		// take user input and clear screen
		fmt.Printf("\n\n\n\n\n\n\t\tPLAYER( %s )>_ ", player)
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
	var s1, s2 int
		_, err := fmt.Sscan(s, &s1, &s2)
		if err != nil {
			fmt.Printf("[Error] Invalid input '%d %d'\n", s1, s2)
			return
		}

		ok, ppoint1 := toCoord(coordinates, s1)
		if !ok {
			continue
		}
		ok, ppoint2 := toCoord(coordinates, s2)
		if !ok {
			continue
		}
		point1 := *ppoint1
		point2 := *ppoint2
		
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
