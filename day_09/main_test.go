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

	t.Run("Knot", func(t *testing.T) {
		t.Run("should record seen positions", func(t *testing.T) {
			knot := NewKnot()

			if got, want := knot.Count(), 1; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			knot.Move(1, 0)
			knot.Move(1, 0)
			knot.Move(0, 1)
			knot.Move(0, 1)

			if got, want := knot.Count(), 5; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			testPosition(t, knot.pos, Position{X: 2, Y: 2})

			knot.Move(0, -1) // already visited

			if got, want := knot.Count(), 5; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			testPosition(t, knot.pos, Position{X: 2, Y: 1})
		})
	})

	t.Run("Rope", func(t *testing.T) {
		t.Run("should move tail to follow head", func(t *testing.T) {
			rope := NewRope(2)

			// move head right 1; tail should not move
			rope.Move(1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 1, Y: 0})
			testPosition(t, rope.knots[1].pos, Position{X: 0, Y: 0})

			// move head right 1; tail should follow
			rope.Move(1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 2, Y: 0})
			testPosition(t, rope.knots[1].pos, Position{X: 1, Y: 0})

			// move head right 1; tail should follow
			rope.Move(1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 3, Y: 0})
			testPosition(t, rope.knots[1].pos, Position{X: 2, Y: 0})

			// move head up 1; tail should not move
			rope.Move(0, 1)
			testPosition(t, rope.knots[0].pos, Position{X: 3, Y: 1})
			testPosition(t, rope.knots[1].pos, Position{X: 2, Y: 0})

			// move head up 1; tail should move diagonally
			rope.Move(0, 1)
			testPosition(t, rope.knots[0].pos, Position{X: 3, Y: 2})
			testPosition(t, rope.knots[1].pos, Position{X: 3, Y: 1})

			// move head up 1; tail should follow
			rope.Move(0, 1)
			testPosition(t, rope.knots[0].pos, Position{X: 3, Y: 3})
			testPosition(t, rope.knots[1].pos, Position{X: 3, Y: 2})

			// move head left 1; tail should not move
			rope.Move(-1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 2, Y: 3})
			testPosition(t, rope.knots[1].pos, Position{X: 3, Y: 2})

			// move head left 1; tail should move diagonally
			rope.Move(-1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 1, Y: 3})
			testPosition(t, rope.knots[1].pos, Position{X: 2, Y: 3})

			// move head left 1; tail should follow
			rope.Move(-1, 0)
			testPosition(t, rope.knots[0].pos, Position{X: 0, Y: 3})
			testPosition(t, rope.knots[1].pos, Position{X: 1, Y: 3})

			// move head down 1; tail should not move
			rope.Move(0, -1)
			testPosition(t, rope.knots[0].pos, Position{X: 0, Y: 2})
			testPosition(t, rope.knots[1].pos, Position{X: 1, Y: 3})

			// move head down 1; tail should move diagonally
			rope.Move(0, -1)
			testPosition(t, rope.knots[0].pos, Position{X: 0, Y: 1})
			testPosition(t, rope.knots[1].pos, Position{X: 0, Y: 2})
		})
	})

	t.Run("Scan", func(t *testing.T) {
		input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"

		t.Run("should parse input to move 2 knot rope", func(t *testing.T) {
			rope := NewRope(2)
			Scan(strings.NewReader(input), rope)

			if got, want := rope.Tail().Count(), 13; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("should parse input to move 10 knot rope", func(t *testing.T) {
			rope := NewRope(10)
			Scan(strings.NewReader(input), rope)

			if got, want := rope.Tail().Count(), 1; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

		t.Run("should parse larger input to move 10 knot rope", func(t *testing.T) {
			long := "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20"
			rope := NewRope(10)
			Scan(strings.NewReader(long), rope)

			if got, want := rope.Tail().Count(), 36; got != want {
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
