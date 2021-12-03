package life_game

import (
	"fmt"
	"io"
	"math/rand"
)

// https://golangify.com/game-life

type Universe [][]bool

func NewUniverse(w, h int) Universe {

	u := make(Universe, h)
	for i := 0; i < h; i++ {
		u[i] = make([]bool, w)
	}
	return u
}

func (u Universe) ShowTo(w io.Writer) {
	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			if u[i][j] {
				_, _ = fmt.Fprint(w, "*")
			} else {
				_, _ = fmt.Fprint(w, " ")
			}
		}
		_, _ = fmt.Fprint(w, "\n")
	}
}

func (u Universe) Seed() {
	max := len(u) * len(u[0])
	countOfLive := max / 4

	for countOfLive > 0 {
		position := rand.Intn(max)
		rowNum := position / len(u[0])
		colNum := position % len(u[0])
		if !u[rowNum][colNum] {
			u[rowNum][colNum] = true
			countOfLive--
		}
	}
}

func (u Universe) liveNeighborsCount(rowNum, colNum int) int {
	left := colNum - 1
	if left < 0 {
		left = 0
	}
	right := colNum + 1
	if right >= len(u[0]) {
		right = len(u[0]) - 1
	}
	up := rowNum - 1
	if up < 0 {
		up = 0
	}
	down := rowNum + 1
	if down >= len(u) {
		down = len(u) - 1
	}

	alive := 0
	for i := up; i <= down; i++ {
		for j := left; j <= right; j++ {
			if i == rowNum && j == colNum {
				continue
			}
			if u[i][j] {
				alive++
			}
		}
	}
	return alive
}

func (u Universe) willBeAlive(rowNum, colNum int) bool {
	neighborsCount := u.liveNeighborsCount(rowNum, colNum)

	if u[rowNum][colNum] {
		if neighborsCount < 2 || neighborsCount > 3 {
			return false
		}
	} else {
		if neighborsCount != 3 {
			return false
		}
	}
	return true
}

func Step(u Universe) Universe {
	newU := NewUniverse(len(u[0]), len(u))
	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			newU[i][j] = u.willBeAlive(i, j)
		}
	}
	return newU
}
