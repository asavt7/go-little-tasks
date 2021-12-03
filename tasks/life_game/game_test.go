package life_game_test

import (
	"bytes"
	"github.com/asavt7/go-little-tasks/tasks/life_game"
	"testing"
)

func TestNewUniverse(t *testing.T) {

	w, h := 10, 15
	u := life_game.NewUniverse(w, h)

	if len(u) != h {
		t.Errorf("err height")
	}
	if len(u[0]) != w {
		t.Errorf("err width")
	}
}

func TestUniverse_ShowTo(t *testing.T) {

	w, h := 2, 2
	u := life_game.NewUniverse(w, h)
	u[0][0] = true
	u[0][1] = true

	var buf bytes.Buffer
	u.ShowTo(&buf)

	if buf.String() != "**\n  \n" {
		t.Errorf("err show method")
	}

}

func TestStep(t *testing.T) {
	t.Run("simple 2-2", func(t *testing.T) {
		w, h := 2, 2
		u := life_game.NewUniverse(w, h)
		u[0][0] = true
		u[0][1] = true

		nu := life_game.Step(u)

		var buf bytes.Buffer
		nu.ShowTo(&buf)

		if buf.String() != "  \n  \n" {
			t.Errorf("err show method")
		}
	})

	t.Run("simple 3-3", func(t *testing.T) {
		w, h := 2, 2
		u := life_game.NewUniverse(w, h)
		u[0][0] = true
		u[0][1] = true

		nu := life_game.Step(u)

		var buf bytes.Buffer
		nu.ShowTo(&buf)

		if buf.String() != "  \n  \n" {
			t.Errorf("err show method")
		}
	})

}
