package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
	timer "github.com/go_pomodoro/timer"
)

func main() {
	totalTime := validateInput()

	tm.Clear()
	tm.MoveCursor(100, 1)
	timer.RunTimer(totalTime)
}

func validateInput() time.Duration {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a number between 1 and 60: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	clockTime, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || (clockTime <= 0 || clockTime > 60) {
		fmt.Println("Oops!")
		validateInput()
	}

	return time.Duration(clockTime)
}
