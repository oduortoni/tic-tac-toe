package main

import (
	t "ttt/ttt"
)

func main() {
	opt := &t.Options{
		TokTop:  "x",
		ColorTop:  t.GetColor("cyan"),
		TokBtm:  "o",
		ColorBtm:  t.GetColor("red"),
		Empty:     ".",
		Dimension: 3,
	}
	t.Play(opt)
}
