package main

import (
	"flag"

	"github.com/Saza-ku/JIS-US-practice/game"
)

var t int

func init() {
	flag.IntVar(&t, "t", 30, "time limit")
	flag.Parse()
}

func main() {
	game.Run(t)
}
