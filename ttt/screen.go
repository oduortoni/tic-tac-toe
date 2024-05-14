package ttt

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	// clear screen
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	if err := clearCmd.Run(); err != nil {
		println(err)
		return
	}
}
