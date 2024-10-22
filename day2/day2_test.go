package day2

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestP1(t *testing.T) {
	input := `5 1 9 5
7 5 3
2 4 6 8`

	want := 18
	got, err := P1(input)
	if err != nil {
		t.Errorf("could not run p1: %v", err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
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
	input := `5 9 2 8
9 4 7 3
3 8 6 5`
	want := 9
	got, err := P2(input)
	if err != nil {
		t.Errorf("could not run p2: %v", err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
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
