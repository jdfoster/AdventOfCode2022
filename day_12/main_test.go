package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"

	t.Run("correctly set start value", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, start, _ := Scan(r)

		got, ok := grid.Height(start)
		if !ok {
			t.Fatal("failed to set start poisiton")
		}

		if want := toNum('a'); got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("correctly set end value", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, _, end := Scan(r)

		got, ok := grid.Height(end)
		if !ok {
			t.Fatal("failed to set end poisiton")
		}

		if want := toNum('z'); got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("finds next positions", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, _, end := Scan(r)

		v := grid.EqualOrAbove(Position{X: 4, Y: 2})

		if got, want := len(v), 1; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		if !end.Equal(v[0]) {
			t.Error("End not found")
		}
	})

	t.Run("calculate shortest path", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, start, end := Scan(r)
		res, ok := Walk(grid, start, end)
		if !ok {
			t.Fatal("failed to find path to goal")
		}

		if got, want := res.Count, 31; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
