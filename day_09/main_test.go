package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Position", func(t *testing.T) {
		t.Run("should add coordinates", func(t *testing.T) {
			origin := Position{X: 0, Y: 0}

			cases := []struct {
				exp  Position
				X, Y int
			}{
				{exp: Position{X: 1, Y: 1}, X: 1, Y: 1},
				{exp: Position{X: -1, Y: -1}, X: -1, Y: -1},
			}

			for _, c := range cases {
				act := origin.Add(c.X, c.Y)
				testPosition(t, act, c.exp)
			}
		})

		t.Run("should calculate difference between positions", func(t *testing.T) {
			origin := Position{X: 0, Y: 0}

			cases := []struct {
				input Position
				exp   Position
			}{
				{
					input: Position{X: 1, Y: 1},
					exp:   Position{X: -1, Y: -1},
				},
				{
					input: Position{X: -1, Y: -1},
					exp:   Position{X: 1, Y: 1},
				},
			}

			for _, c := range cases {
				act := origin.Diff(c.input)
				testPosition(t, act, c.exp)
			}
		})
	})

	t.Run("Ending", func(t *testing.T) {
		t.Run("should record seen positions", func(t *testing.T) {
			ending := NewEnding()

			if got, want := ending.Count(), 1; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			ending.Move(1, 0)
			ending.Move(1, 0)
			ending.Move(0, 1)
			ending.Move(0, 1)

			if got, want := ending.Count(), 5; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			testPosition(t, ending.pos, Position{X: 2, Y: 2})

			ending.Move(0, -1) // already visited

			if got, want := ending.Count(), 5; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			testPosition(t, ending.pos, Position{X: 2, Y: 1})
		})
	})

	t.Run("Rope", func(t *testing.T) {
		t.Run("should move tail to follow head", func(t *testing.T) {
			rope := NewRope()

			// move head right 1; tail should not move
			rope.Move(1, 0)
			testPosition(t, rope.Head.pos, Position{X: 1, Y: 0})
			testPosition(t, rope.Tail.pos, Position{X: 0, Y: 0})

			// move head right 1; tail should follow
			rope.Move(1, 0)
			testPosition(t, rope.Head.pos, Position{X: 2, Y: 0})
			testPosition(t, rope.Tail.pos, Position{X: 1, Y: 0})

			// move head right 1; tail should follow
			rope.Move(1, 0)
			testPosition(t, rope.Head.pos, Position{X: 3, Y: 0})
			testPosition(t, rope.Tail.pos, Position{X: 2, Y: 0})

			// move head up 1; tail should not move
			rope.Move(0, 1)
			testPosition(t, rope.Head.pos, Position{X: 3, Y: 1})
			testPosition(t, rope.Tail.pos, Position{X: 2, Y: 0})

			// move head up 1; tail should move diagonally
			rope.Move(0, 1)
			testPosition(t, rope.Head.pos, Position{X: 3, Y: 2})
			testPosition(t, rope.Tail.pos, Position{X: 3, Y: 1})

			// move head up 1; tail should follow
			rope.Move(0, 1)
			testPosition(t, rope.Head.pos, Position{X: 3, Y: 3})
			testPosition(t, rope.Tail.pos, Position{X: 3, Y: 2})

			// move head left 1; tail should not move
			rope.Move(-1, 0)
			testPosition(t, rope.Head.pos, Position{X: 2, Y: 3})
			testPosition(t, rope.Tail.pos, Position{X: 3, Y: 2})

			// move head left 1; tail should move diagonally
			rope.Move(-1, 0)
			testPosition(t, rope.Head.pos, Position{X: 1, Y: 3})
			testPosition(t, rope.Tail.pos, Position{X: 2, Y: 3})

			// move head left 1; tail should follow
			rope.Move(-1, 0)
			testPosition(t, rope.Head.pos, Position{X: 0, Y: 3})
			testPosition(t, rope.Tail.pos, Position{X: 1, Y: 3})

			// move head down 1; tail should not move
			rope.Move(0, -1)
			testPosition(t, rope.Head.pos, Position{X: 0, Y: 2})
			testPosition(t, rope.Tail.pos, Position{X: 1, Y: 3})

			// move head down 1; tail should move diagonally
			rope.Move(0, -1)
			testPosition(t, rope.Head.pos, Position{X: 0, Y: 1})
			testPosition(t, rope.Tail.pos, Position{X: 0, Y: 2})
		})
	})

	t.Run("Scan", func(t *testing.T) {
		input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"

		t.Run("should parse input", func(t *testing.T) {
			r := strings.NewReader(input)
			rope := Scan(r)

			if got, want := rope.Tail.Count(), 13; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}

func testPosition(t *testing.T, actual, expected Position) {
	t.Helper()

	if got, want := actual.X, expected.X; got != want {
		t.Errorf("x; got %d, want %d", got, want)
	}

	if got, want := actual.Y, expected.Y; got != want {
		t.Errorf("y; got %d, want %d", got, want)
	}
}
