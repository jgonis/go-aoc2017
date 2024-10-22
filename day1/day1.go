package day1

import (
	"strconv"
)

func P1(input string) (int, error) {
	digits, err := stringToNumSlice(input)
	if err != nil {
		return 0, err
	}

	sum := findSum(digits, 1)
	return sum, nil
}

func P2(input string) (int, error) {
	digits, err := stringToNumSlice(input)
	if err != nil {
		return 0, err
	}
	lookahead := len(digits) / 2
	sum := findSum(digits, lookahead)

	return sum, nil
}

func stringToNumSlice(input string) ([]int, error) {
	digits := make([]int, 0)
	for _, digit_rune := range input {
		digit, err := strconv.Atoi(string(digit_rune))
		if err != nil {
			return nil, err
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func findSum(digits []int, lookahead int) int {
	sum := 0
	for index, digit := range digits {
		nextIndex := (index + lookahead) % len(digits)
		if digit == digits[nextIndex] {
			sum += digit
		}
	}
	return sum
}
