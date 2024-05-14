package main

import (
	t "ttt/ttt"
)

func main() {
	opt := &t.Options{
		Token1: "x",
		Token2: "o",
		Empty: ".",
		Dimension: 3,
	}
	t.Play(opt)
}
