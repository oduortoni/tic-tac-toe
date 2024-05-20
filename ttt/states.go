package ttt

import (
	"fmt"
)

func print_exit(player string) {
	print(RED)
	fmt.Printf("\n\n\n\n\t\t[ Player ")
	print(CYAN)
	fmt.Printf("%s", player)
	print(RED)
	fmt.Printf(" exited ]")
	print(ORIGINAL)
	fmt.Printf("\n\n\n\n\n\n")
}

func print_success(grid Grid, options *Options, player string) {
	ClearScreen()
	success_msg := fmt.Sprintf("\n\n\n\n\t\t[ Player %q has won the game ]", player)
	println(GREEN + success_msg + ORIGINAL)
	println()
	fmt.Printf("\n\n\t\t...")
	print_grid(grid, options)
	print("\n\n\n\n\n")
}
