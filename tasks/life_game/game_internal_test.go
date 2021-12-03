package life_game

import "testing"

func Test_liveNeighborsCount(t *testing.T) {
	w, h := 2, 2
	u := NewUniverse(w, h)
	u[0][0] = true
	u[0][1] = true

	if actual := u.liveNeighborsCount(0, 0); actual != 1 {
		t.Errorf("expected %d, actual %d", 1, actual)
	}

	if actual := u.liveNeighborsCount(1, 1); actual != 2 {
		t.Errorf("expected %d, actual %d", 2, actual)

	}
}
