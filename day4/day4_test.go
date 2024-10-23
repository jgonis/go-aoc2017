package day4

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "aa bb cc dd ee",
			want:  1,
		},
		{
			input: "aa bb cc dd aa",
			want:  0,
		},
		{
			input: "aa bb cc dd aaa",
			want:  1,
		},
		{
			input: `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`,
			want: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test case input:  "+testCase.input, func(t *testing.T) {
			got := P1(testCase.input)
			if testCase.want != got {
				t.Errorf("want %v, got %v", testCase.want, got)
			}
		})
	}
}

func TestP1RealInput(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Errorf("could not read input file: %v", err)
	}
	trimmedInput := strings.TrimSpace(string(input))
	validCount := P1(trimmedInput)
	fmt.Println(validCount)
}

func TestP2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "abcde fghij",
			want:  1,
		},
		{
			input: "abcde xyz ecdab",
			want:  0,
		},
		{
			input: "a ab abc abd abf abj",
			want:  1,
		},
		{
			input: "iiii oiii ooii oooi oooo",
			want:  1,
		},
		{
			input: "oiii ioii iioi iiio",
			want:  0,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test case input:  "+testCase.input, func(t *testing.T) {
			got := P2(testCase.input)
			if testCase.want != got {
				t.Errorf("want %v, got %v", testCase.want, got)
			}
		})
	}
}

func TestP2RealInput(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Errorf("could not read input file: %v", err)
	}
	trimmedInput := strings.TrimSpace(string(input))
	validCount := P2(trimmedInput)
	fmt.Println(validCount)
}
