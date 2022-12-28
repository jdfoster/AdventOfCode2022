package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Scan", func(t *testing.T) {
		t.Run("should parse input", func(t *testing.T) {
			cases := []Location{
				{
					Sensor: Position{X: 2, Y: 18},
					Beacon: Position{X: -2, Y: 15},
				},
				{
					Sensor: Position{X: 9, Y: 16},
					Beacon: Position{X: 10, Y: 16},
				},
				{
					Sensor: Position{X: 13, Y: 2},
					Beacon: Position{X: 15, Y: 3},
				},
				{
					Sensor: Position{X: 12, Y: 14},
					Beacon: Position{X: 10, Y: 16},
				},
			}

			r := strings.NewReader(input)
			locs := Scan(r)

			if got, want := len(locs), 14; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			for i, c := range cases {
				ContrastLocation(t, locs[i], c)
			}
		})
	})

	t.Run("Position", func(t *testing.T) {
		origin := Position{X: 0, Y: 0}
		cases := []struct {
			p  Position
			d  int
			ll int
		}{
			{
				p:  Position{X: 6, Y: 6},
				d:  12,
				ll: 311,
			},
			{
				p:  Position{X: 3, Y: 3},
				d:  6,
				ll: 83,
			},
			{
				p:  Position{X: 2, Y: 2},
				d:  4,
				ll: 39,
			},
		}

		t.Run("distance should calculate manhattan distance", func(t *testing.T) {
			for _, c := range cases {
				if got, want := origin.Distance(c.p), c.d; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			}
		})

		t.Run("dilate should list point within distance", func(t *testing.T) {
			for _, c := range cases {
				ll := origin.Dilate(c.p)

				if got, want := len(ll), c.ll; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			}
		})
	})

	t.Run("Radar", func(t *testing.T) {
		t.Run("should build and count characters", func(t *testing.T) {
			r := strings.NewReader(input)
			locs := Scan(r)
			radar := NewRadar(locs)

			if got, want := radar.CountRow(10, '#'), 26; got != want {
				fmt.Println(radar)
				t.Errorf("got %d, want %d", got, want)
			}
		})
	})
}

func ContrastLocation(t *testing.T, a, b Location) {
	t.Helper()

	ContrastPosition(t, a.Sensor, b.Sensor)
	ContrastPosition(t, a.Beacon, b.Beacon)
}

func ContrastPosition(t *testing.T, a, b Position) {
	t.Helper()

	if got, want := a.X, b.X; got != want {
		t.Errorf("X; got %d, want %d", got, want)
	}

	if got, want := a.Y, b.Y; got != want {
		t.Errorf("X; got %d, want %d", got, want)
	}
}

var input = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
