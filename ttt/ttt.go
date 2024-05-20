package ttt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Play(options *Options) {
	var turn bool = true
	// process input
	var point1, point2 Point
	computerTok := ""
	coordinates := toCoordinateSystem(options.Dimension)
	grid := init_grid(options.SymbolTop, options.SymbolBtm, options.Empty, options.Dimension)
	reader := bufio.NewReader(os.Stdout)

	// assign computer a token
	if options.Computer {
		computerTok = options.SymbolComp
	}


	// begin with the initial layout
	ClearScreen()
	print_grid(grid, options)

	// which player begins
	player := options.Starter
	for {
		if options.Computer && turn {
			turn = false


			p1, p2 := computeCoords(grid, options.Empty, computerTok)
			point1 = *p1
			point2 = *p2
			// fmt.Printf("(%d, %d) to (%d, %d)\n", point1.X, point1.Y, point2.X, point2.Y)
			// return
		} else {
			turn = true

			// take user input and clear screen
			fmt.Printf("\n\n\n\n\n\n\t\tPLAYER( %s )>_ ", player)
			s, _ := reader.ReadString(10)
			s = strings.TrimSpace(s)

			// quit game
			if s == "quit" || s == "q" || s == "exit" || s == "e" {
				print_exit(player)
				return
			}

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
			point1 = *ppoint1
			point2 = *ppoint2
		}


		// clear and update screen each iteration
		ClearScreen()
		// before update
		print_grid(grid, options)
		
		if player == options.SymbolTop {
			if updated := update_grid(grid, point1, point2, options.Empty, options.SymbolTop); updated {
				player = options.SymbolBtm
			}
		} else {
			if updated := update_grid(grid, point1, point2, options.Empty, options.SymbolBtm); updated {
				player = options.SymbolTop
			}
		}
		// screen after update
		print_grid(grid, options)

		// check for winner
		if found, who := winner(grid, options); found {
			print_success(grid, options, who)
			return
		}
	}
}
