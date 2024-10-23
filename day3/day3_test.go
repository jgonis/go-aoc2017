package day3

import (
	"fmt"
	"strconv"
	"testing"
)

func TestP1(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			input: 1,
			want:  0,
		},
		{
			input: 12,
			want:  3,
		},
		{
			input: 23,
			want:  2,
		},
		{
			input: 1024,
			want:  31,
		},
	}
	for _, testCase := range testCases {
		t.Run("Test case for value: "+strconv.Itoa(testCase.input), func(t *testing.T) {
			got := P1(testCase.input)
			if got != testCase.want {
				t.Errorf("want %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestP1RealInput(t *testing.T) {
	output := P1(325489)
	fmt.Println(output)
}

func TestP2(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			input: 4,
			want:  5,
		},
		{
			input: 11,
			want:  23,
		},
		{
			input: 26,
			want:  54,
		},
		{
			input: 147,
			want:  304,
		},
	}
	for _, testCase := range testCases {
		t.Run("Test case for value: "+strconv.Itoa(testCase.input), func(t *testing.T) {
			got := P2(testCase.input)
			if got != testCase.want {
				t.Errorf("want %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestP2RealInput(t *testing.T) {
	output := P2(325489)
	fmt.Println(output)
}
