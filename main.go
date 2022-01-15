package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Saza-ku/JIS-US-practice/game"
)

var t int

func init() {
	flag.IntVar(&t, "t", 30, "time limit")
	flag.Parse()
}

func main() {
	if t <= 0 {
		fmt.Println("time limit must be more than 0 seconds")
		os.Exit(1)
	}
	game.Run(t)
}
