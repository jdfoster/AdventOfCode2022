package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("should parse input", func(t *testing.T) {
		r := strings.NewReader(input)
		res := Scan(r)

		if got, want := len(res), len(details); got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		for _, want := range details {
			got := res[want.Name]
			CompaeValueMeta(t, got, want)
		}
	})

	t.Run("should build root", func(t *testing.T) {
		r := strings.NewReader(input)
		d := Scan(r)
		vv := d.Build("AA")

		test := []string{"DD", "II", "BB"}

		for i, v := range vv.Tunnel {
			tc := test[i]

			if got, want := v.Name, tc; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})

	t.Run("should record open values", func(t *testing.T) {
		r := strings.NewReader(input)
		d := Scan(r)
		w := &Walker{Current: d.Build("AA")}

		mv := func(n string) {
			if !w.Move(n) {
				t.Errorf("failed to move to value %q", n)
			}
		}

		isp := func(rate int, open bool) {
			r, o := w.Inspect()

			if got, want := r, rate; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			if got, want := o, open; got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		}

		isp(0, true)

		mv("DD")
		isp(20, true)
		w.Open()
		isp(20, false)

		mv("CC")
		isp(2, true)
		w.Open()
		isp(2, false)
	})

	t.Run("should walk to find maximum flow", func(t *testing.T) {
		r := strings.NewReader(input)
		d := Scan(r)
		g := d.Build("AA")

		if got, want := Walk(g, 30), 1651; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func CompaeValueMeta(t *testing.T, a, b ValueDetail) {
	t.Helper()

	if got, want := a.Name, b.Name; got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if got, want := a.FlowRate, b.FlowRate; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	if got, want := len(a.Values), len(b.Values); got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	for i, want := range b.Values {
		if got := a.Values[i]; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

var details = []ValueDetail{
	{
		Name:     "AA",
		FlowRate: 0,
		Values:   []string{"DD", "II", "BB"},
	},
	{
		Name:     "BB",
		FlowRate: 13,
		Values:   []string{"CC", "AA"},
	},
	{
		Name:     "CC",
		FlowRate: 2,
		Values:   []string{"DD", "BB"},
	},
	{
		Name:     "DD",
		FlowRate: 20,
		Values:   []string{"CC", "AA", "EE"},
	},
	{
		Name:     "EE",
		FlowRate: 3,
		Values:   []string{"FF", "DD"},
	},
	{
		Name:     "FF",
		FlowRate: 0,
		Values:   []string{"EE", "GG"},
	},
	{
		Name:     "GG",
		FlowRate: 0,
		Values:   []string{"FF", "HH"},
	},
	{
		Name:     "HH",
		FlowRate: 22,
		Values:   []string{"GG"},
	},
	{
		Name:     "II",
		FlowRate: 0,
		Values:   []string{"AA", "JJ"},
	},
	{
		Name:     "JJ",
		FlowRate: 21,
		Values:   []string{"II"},
	},
}

var input = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`
