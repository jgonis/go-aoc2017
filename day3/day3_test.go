package day3

import "testing"

func TestP1(t *testing.T) {
	input := `17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...`
	want := 0
	got := P1(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
