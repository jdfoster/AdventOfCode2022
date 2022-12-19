package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := "30373\n25512\n65332\n33549\n35390"

	t.Run("Trees collection", func(t *testing.T) {
		t.Run("should count tree without duplicate", func(t *testing.T) {
			tt := Trees{
				{X: 0, Y: 0},
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			}

			set := make(TreeSet)
			set.AddTree(tt...)

			if got, want := set.Count(), 2; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("should count visable trees", func(t *testing.T) {
			tt := Trees{
				{Y: 0, X: 0, H: 3},
				{Y: 0, X: 1, H: 0},
				{Y: 0, X: 2, H: 7},
				{Y: 0, X: 3, H: 3},
			}

			res := tt.Visible()
			if got, want := res.Count(), 3; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})

	t.Run("Parse tree input", func(t *testing.T) {
		r := strings.NewReader(input)
		grid := Scan(r)

		if got, want := len(grid.rows), 5; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		for _, row := range grid.rows {
			if got, want := len(row), 5; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		}

		expected := []int{3, 2, 6, 3, 3}

		var actual []int
		for _, item := range grid.Column(0) {
			actual = append(actual, item.H)
		}

		if got, want := len(expected), len(actual); got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		for i, want := range expected {
			if got := actual[i]; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		}
	})

	t.Run("Count visible trees", func(t *testing.T) {
		r := strings.NewReader(input)
		grid := Scan(r)
		set := grid.Visible()

		if got, want := set.Count(), 21; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
