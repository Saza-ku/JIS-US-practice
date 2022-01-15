package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		r := bufio.NewReaderSize(os.Stdin, 1000000)
		x := 0
		for {
			s, _, err := r.ReadLine()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(s))
			x++
			ch <- x
		}
	}()
	x := 0
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("DONE")
			fmt.Println(x)
			return
		case x = <-ch:
		}
	}
}
