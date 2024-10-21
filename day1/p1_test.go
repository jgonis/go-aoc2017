package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "1122",
			want:  3,
		},
		{
			input: "1111",
			want:  4,
		},
		{
			input: "1234",
			want:  0,
		},
		{
			input: "91212129",
			want:  9,
		},
	}
	for caseNumber, testCase := range testCases {
		t.Run("Test case "+strconv.Itoa(caseNumber), func(t *testing.T) {
			got, _ := P1(testCase.input)
			if got != testCase.want {
				t.Errorf("want %d, got %d", testCase.want, got)
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
	output, err := P1(trimmedInput)
	if err != nil {
		t.Errorf("could not run p1: %v", err)
	}
	fmt.Println(output)
}

func TestP2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "1212",
			want:  6,
		},
		{
			input: "1221",
			want:  0,
		},
		{
			input: "123425",
			want:  4,
		},
		{
			input: "123123",
			want:  12,
		},
		{
			input: "12131415",
			want:  4,
		},
	}

	for caseNumber, testCase := range testCases {
		t.Run("Test case "+strconv.Itoa(caseNumber), func(t *testing.T) {
			got, _ := P2(testCase.input)
			if got != testCase.want {
				t.Errorf("given: %v, want %d, got %d", testCase.input, testCase.want, got)
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
	output, err := P2(trimmedInput)
	if err != nil {
		t.Errorf("could not run p1: %v", err)
	}
	fmt.Println(output)
}

func TestStringToNumSlice(t *testing.T) {
	testCases := []struct {
		input     string
		want      []int
		wantError bool
	}{
		{
			input:     "",
			want:      []int{},
			wantError: false,
		},
		{
			input:     "1",
			want:      []int{1},
			wantError: false,
		},
		{
			input:     "123",
			want:      []int{1, 2, 3},
			wantError: false,
		},
		{
			input:     "123a",
			want:      nil,
			wantError: true,
		},
	}
	for caseNumber, testCase := range testCases {
		t.Run("Test case "+strconv.Itoa(caseNumber), func(t *testing.T) {
			got, err := stringToNumSlice(testCase.input)
			if err != nil && !testCase.wantError {
				t.Errorf("want no error, got %v", err)
			}
			if err == nil && testCase.wantError {
				t.Errorf("want error, got no error")
			}
			if !slices.Equal(got, testCase.want) {
				t.Errorf("want %v, got %v", testCase.want, got)
			}
		})
	}
}
