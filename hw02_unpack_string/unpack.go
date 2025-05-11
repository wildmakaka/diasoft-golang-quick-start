package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {

	var res string = ""

	// if res != "" {

	// 	errors.New("invalid string")

	// }

	if ErrInvalidString != nil {
		return "wtf", ErrInvalidString
	}

	res = firstStep(input)
	res = secondStep(reverseString(res))
	res = reverseString(res)

	fmt.Printf("\n result ")
	fmt.Printf(res)

	fmt.Printf("\n")
	// return "", nil
	return res, nil
}

func firstStep(input string) string {

	var res string = ""
	for index, runeValue := range input {

		if isNumeric(string(runeValue)) {
			var prevSymbol string = string(input[index-1])
			var repeatTimes int = runeToNumber(runeValue) - 1

			if runeToNumber(runeValue) > 0 {
				repeatTimes = runeToNumber(runeValue) - 1
				res = res + repeatSymbol(prevSymbol, repeatTimes)
			} else {
				res = res + "0"
			}

		} else {
			res = res + string(runeValue)
		}
	}
	return res
}

func secondStep(input string) string {

	var res string = ""
	var skipNextSymbol bool = false

	for _, runeValue := range input {

		// fmt.Printf("isTrue: %t\n", skipNextSymbol)

		if skipNextSymbol == false {
			if isNumeric(string(runeValue)) {
				skipNextSymbol = true
				continue
			} else {
				skipNextSymbol = false
				res = res + string(runeValue)
			}
		} else {
			skipNextSymbol = false
		}
	}
	return res
}

func isNumeric(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func repeatSymbol(symbol string, repeatTimes int) string {
	return strings.Repeat(symbol, repeatTimes)
}

func runeToNumber(inputRune rune) int {
	if num, err := strconv.Atoi(string(inputRune)); err == nil {
		return num
	}
	return 0
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
