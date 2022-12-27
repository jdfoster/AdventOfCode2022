package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"

	t.Run("Rocks", func(t *testing.T) {
		t.Run("should find boundary", func(t *testing.T) {
			r := strings.NewReader(input)
			rocks := Scan(r)
			a, b, c, d := rocks.FindBoundry()
			vs := []int{a, b, c, d}

			cases := []int{494, 503, 4, 9}

			for i, want := range cases {
				if got := vs[i]; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			}
		})
	})

	t.Run("Position", func(t *testing.T) {
		t.Run("should list points between two positions", func(t *testing.T) {
			cases := []struct {
				a Position
				b Position
				p []Position
			}{
				{
					a: Position{X: 498, Y: 4},
					b: Position{X: 498, Y: 6},
					p: []Position{{X: 498, Y: 4}, {X: 498, Y: 5}, {X: 498, Y: 6}},
				},
				{
					a: Position{X: 498, Y: 6},
					b: Position{X: 496, Y: 6},
					p: []Position{{X: 498, Y: 6}, {X: 497, Y: 6}, {X: 496, Y: 6}},
				},
			}

			for _, c := range cases {
				v := c.a.PointsBetween(c.b)
				if got, want := len(v), len(c.p); got != want {
					t.Fatalf("got %d, want %d", got, want)
				}

				for i, want := range c.p {
					if got := v[i]; got != want {
						t.Errorf("got %d, want %d", got, want)
					}
				}
			}
		})
	})

	t.Run("Cave", func(t *testing.T) {
		t.Run("should draw rock formations", func(t *testing.T) {
			want := ".......+...\n...........\n...........\n...........\n.....#...##\n.....#...#.\n...###...#.\n.........#.\n.........#.\n.#########."

			r := strings.NewReader(input)
			rocks := Scan(r)
			cave := NewCave(rocks)

			if got := cave.String(); strings.TrimSpace(got) != strings.TrimSpace(want) {
				t.Errorf("got  >>>\n%s", got)
				t.Errorf("want >>>\n%s", want)
			}
		})

		t.Run("should add sand till full", func(t *testing.T) {
			want := ".......+...\n...........\n.......o...\n......ooo..\n.....#ooo##\n....o#ooo#.\n...###ooo#.\n.....oooo#.\n..o.ooooo#.\n.#########."

			r := strings.NewReader(input)
			rocks := Scan(r)
			cave := NewCave(rocks)

			var count int
			for i := 0; i < 100; i++ {
				if cave.AddSand() {
					count++
				}
			}

			if got, want := count, 24; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			if got, want := cave.Count('o'), 24; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			if got := cave.String(); strings.TrimSpace(got) != strings.TrimSpace(want) {
				t.Errorf("got  >>>\n%s", got)
				t.Errorf("want >>>\n%s", want)
			}
		})
	})
}
