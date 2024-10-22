package day2

import (
	"slices"
	"strconv"
	"strings"
)

func P1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	checksum := 0
	for _, line := range lines {
		digits, err := lineToSortedNumberSlice(line)
		if err != nil {
			return 0, err
		}

		lineChecksum := digits[len(digits)-1] - digits[0]
		checksum += lineChecksum
	}
	return checksum, nil
}

func P2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	checksum := 0
	for _, line := range lines {
		digits, err := lineToSortedNumberSlice(line)
		if err != nil {
			return 0, err
		}
		a, b := findDivisibleNumbers(digits)
		checksum += b / a

	}
	return checksum, nil
}

func findDivisibleNumbers(digits []int) (int, int) {
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[j]%digits[i] == 0 {
				return digits[i], digits[j]
			}
		}
	}
	return 0, 0
}

func lineToSortedNumberSlice(line string) ([]int, error) {
	line = strings.TrimSpace(line)
	digitStrings := strings.Fields(line)
	digits := make([]int, 0)
	for _, digitString := range digitStrings {
		digitString := strings.TrimSpace(digitString)
		digit, err := strconv.Atoi(digitString)
		if err != nil {
			return nil, err
		}
		digits = append(digits, digit)
	}
	slices.Sort(digits)
	return digits, nil
}
