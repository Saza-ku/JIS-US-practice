package game

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type result int

const (
	success result = iota
	fail
)

func Run() {
	ch := start()

	correct := 0
	incorrect := 0
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println()
			fmt.Println("====== TIME OUT ======")
			fmt.Println("RESULT:")
			fmt.Printf("  CORRECT  : %d\n", correct)
			fmt.Printf("  INCORRECT: %d\n", incorrect)
			return
		case r := <-ch:
			if r == success {
				correct++
			} else {
				incorrect++
			}
		}
	}
}

func start() <-chan result {
	ch := make(chan result)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			question := newTask()
			fmt.Printf("TASK: %s", question)
			fmt.Println()
			ans, _, err := reader.ReadLine()
			if err != nil {
				panic(err)
			}
			if question == string(ans) {
				ch <- success
				fmt.Println("CORRECT!")
			} else {
				ch <- fail
				fmt.Println("INCORRECT!")
			}
			fmt.Println()
		}
	}()
	return ch
}

func newTask() string {
	symbols := []string{"^", "&", "*", "(", ")", "-", "_", "=", "+", "¥", "|", "\\", "[", "{", "]", "}", ";", ":", "\"", "'", "`", "~"}
	task := ""

	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := (rand.Intn(len(symbols)))
		task += symbols[idx]
	}

	return task
}
