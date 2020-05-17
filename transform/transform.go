package transform

import (
	tm "github.com/buger/goterm"
)

// ExtractNumbers : returns the ten and one place ints in a two digit number
func ExtractNumbers(twoDigitNumber int) (int, int) {

	tenDigit := (twoDigitNumber / 10) % 10
	onesDigit := (twoDigitNumber % 10)

	return tenDigit, onesDigit
}

// FormatNumber : prints the clock digits
func FormatNumber(minutesTens, minuteOnes, place, secondTens, secondOnes [5]string) {
	for i := 0; i < len(minutesTens); i++ {
		tm.MoveCursorForward(60)
		tm.Println(minutesTens[i], minuteOnes[i], place[i], secondTens[i], secondOnes[i])
	}
}
