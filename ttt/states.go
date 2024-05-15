package ttt

import (
	"fmt"
)

func print_exit(grid Grid, player string) {
	ClearScreen()
	fmt.Printf("\n\n\n\n\t\t[ Player %q exited ]", player)
	fmt.Printf("\n\n\t\tThe final board configuration is...")
}

func print_success(grid Grid, player string) {
	ClearScreen()
	success_msg := fmt.Sprintf("\n\n\n\n\t\t[ Player %q has won the game ]", player)
	println(GREEN + success_msg + ORIGINAL)
	println()
	fmt.Printf("\n\n\t\tThe final board configuration is...")
}
