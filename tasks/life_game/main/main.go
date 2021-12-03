package main

import (
	"fmt"
	"github.com/asavt7/go-little-tasks/tasks/life_game"
	"os"
	"os/exec"
	"time"
)

func RunSimulator(w, h int) {
	step := 0

	clearScreen()
	u := life_game.NewUniverse(w, h)
	u.Seed()

	for {
		clearScreen()
		_, _ = fmt.Fprintf(os.Stdout, "step : %d\n", step)
		u.ShowTo(os.Stdout)
		u = life_game.Step(u)
		step++
		time.Sleep(1500 * time.Millisecond)

	}
}

func main() {
	RunSimulator(50, 15)
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
