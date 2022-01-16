package game

import (
	"bufio"
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

func Run(t int) {
	ch := start()
	fmt.Println("====== START ======")

	correct := 0
	incorrect := 0
	timeout := time.After(time.Duration(t) * time.Second)

	for {
		select {
		case <-timeout:
			showResult(correct, incorrect)
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

	go questionLoop(ch)

	return ch
}

func questionLoop(ch chan<- result) {
	reader := bufio.NewReader(os.Stdin)
	for {
		question := newTask()
		fmt.Printf("TASK: %s\n", question)

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
}

func newTask() string {
	symbols := []string{"^", "&", "*", "(", ")", "-", "_", "=", "+", "|", "\\", "[", "{", "]", "}", ";", ":", "\"", "'", "`", "~"}
	task := ""

	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := (rand.Intn(len(symbols)))
		task += symbols[idx]
	}

	return task
}

func showResult(correct, incorrect int) {
	fmt.Println()
	fmt.Println("====== TIME OUT ======")
	fmt.Println("RESULT:")
	fmt.Printf("  CORRECT  : %d\n", correct)
	fmt.Printf("  INCORRECT: %d\n", incorrect)
}
