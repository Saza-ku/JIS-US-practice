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
		r := bufio.NewReader(os.Stdin)
		point := 0
		for {
			question := "kyoto"
			fmt.Printf("TASK: %s", question)
			fmt.Println()
			ans, _, err := r.ReadLine()
			if err != nil {
				panic(err)
			}
			if question == string(ans) {
				point++
			} else {
				fmt.Println("incorrect!")
			}
			ch <- point
			fmt.Println()
		}
	}()

	point := 0
	fmt.Println(time.Now())
	for {
		select {
		case t := <-time.After(5 * time.Second):
			fmt.Println()
			fmt.Println("====== TIME OUT ======")
			fmt.Printf("Your point is %d!\n", point)
			fmt.Println(t)
			return
		case point = <-ch:
		}
	}
}
