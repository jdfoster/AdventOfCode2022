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
		next := Position{X: 4, Y: 2}

		v := grid.EqualOrLower(end)

		if got, want := len(v), 1; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		if !next.Equal(v[0]) {
			t.Error("end not found")
		}
	})

	t.Run("calculate shortest path", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, start, end := Scan(r)
		res, ok := Walk(grid, end, start.Equal)
		if !ok {
			t.Fatal("failed to find path to goal")
		}

		if got, want := res.Count, 31; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("calculate closest square with elevation \"a\"", func(t *testing.T) {
		r := strings.NewReader(input)
		grid, _, end := Scan(r)
		fin := grid.FindElevation('a')
		res, ok := Walk(grid, end, fin)
		if !ok {
			t.Fatal("failed to find square with elevation 'a'")
		}

		if got, want := res.Count, 29; got != want {
			t.Errorf("got %d, wamt %d", got, want)
		}
	})
}
