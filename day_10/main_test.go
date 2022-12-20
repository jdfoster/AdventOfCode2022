package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	cases := []struct {
		k int
		v int
	}{
		{k: 20, v: 21},
		{k: 60, v: 19},
		{k: 100, v: 18},
		{k: 140, v: 21},
		{k: 180, v: 16},
		{k: 220, v: 18},
	}

	t.Run("Calculate strength", func(t *testing.T) {
		proc := NewProc()

		if got, want := len(proc.state), 0; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		if got, want := proc.Strength(), 0; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		r := strings.NewReader(input)
		insts := Scan(r)
		mach := NewMachine()
		proc.Run(mach, insts)

		if got, want := len(proc.state), 6; got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				mem, ok := proc.state[c.k]
				if !ok {
					t.Fatalf("key %d not found", c.k)
				}

				if got, want := mem.x, c.v; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			})
		}

		if got, want := proc.Strength(), 13140; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Render output", func(t *testing.T) {
		r := strings.NewReader(input)
		items := Scan(r)
		mach := NewMachine()
		proc := NewProc()
		proc.Run(mach, items)

		got := strings.TrimSpace(proc.crt.Render())
		want := strings.TrimSpace(output)

		if got != want {
			t.Error("failed to draw outout")
			fmt.Println("got >>>")
			fmt.Println(got)
			fmt.Println("want >>>")
			fmt.Println(want)
		}
	})
}

var input = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

var output = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
