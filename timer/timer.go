package timer

import (
	"strconv"
	"time"

	tm "github.com/buger/goterm"
	"github.com/go_pomodoro/digits"
	"github.com/go_pomodoro/transform"
)

const current = "Current Time"
const finish = "Finish Time"

// RunTimer : Maintains current timer
func RunTimer(totalTime time.Duration) {
	second := 59

	startTime := CurrentTime()
	finishTime := startTime.Add(time.Minute * totalTime)
	print(finish, finishTime)
	progress := 0

	for {
		curTime := CurrentTime()

		tm.MoveCursor(1, 1)
		print(current, curTime)

		if breakLoop(curTime, finishTime) {
			break
		}

		minuteTens, minuteOnes := transform.ExtractNumbers(int(time.Until(finishTime).Minutes()))
		secondTens, secondOnes := transform.ExtractNumbers(second)

		transform.FormatNumber(digits.GetDigits(strconv.Itoa(minuteTens)),
			digits.GetDigits(strconv.Itoa(minuteOnes)),
			digits.GetDigits(":"),
			digits.GetDigits(strconv.Itoa(secondTens)),
			digits.GetDigits(strconv.Itoa(secondOnes)))

		tm.Flush()
		time.Sleep(time.Second)
		second = maintainSeconds(second)

		progress--
	}
}

// CurrentTime : returns the current timer
func CurrentTime() time.Time {
	return time.Now()
}

func breakLoop(curTime time.Time, finishTime time.Time) bool {
	if curTime.Format(time.RFC1123) > finishTime.Format(time.RFC1123) {
		return true
	}
	return false
}

func maintainSeconds(second int) int {
	second--
	if second <= 0 {
		second = 59
	}
	return second
}

func print(s string, curTime time.Time) {
	tm.Println(s, curTime.Format(time.RFC1123))
}
