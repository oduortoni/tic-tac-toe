package main

import (
	t "ttt/ttt"
)

func main() {
	opt := &t.Options{
		Token1: "x",
		Color1: t.GetColor("cyan"),
		Token2: "o",
		Color2: t.GetColor("red"),
		Empty: ".",
		Dimension: 3,
	}
	t.Play(opt)
}
