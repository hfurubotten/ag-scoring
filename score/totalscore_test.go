package score

import "testing"

var scores = []struct {
	in  []*Score
	out int
}{
	{[]*Score{setScore(10, 10, 1), setScore(5, 5, 1), setScore(15, 15, 1)}, 100},
	{[]*Score{setScore(5, 10, 1), setScore(5, 5, 1), setScore(20, 40, 1)}, 66},
	{[]*Score{setScore(5, 10, 1), setScore(5, 10, 1), setScore(20, 40, 1)}, 50},
	{[]*Score{setScore(10, 10, 2), setScore(5, 10, 1), setScore(20, 40, 1)}, 75},
}

func setScore(points, max, w int) *Score {
	s := NewScore(max, w)
	s.IncBy(points)
	return s
}

func TestTotal(t *testing.T) {
	for _, s := range scores {
		tot := Total(s.in)
		if tot != s.out {
			t.Errorf("Got: %d, Want: %d", tot, s.out)
		}
	}
}
