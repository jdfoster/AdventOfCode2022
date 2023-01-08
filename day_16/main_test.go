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
			got := res[want.ID]
			CompaeValueMeta(t, got, want)
		}
	})

	t.Run("should walk to find maximum flow", func(t *testing.T) {
		r := strings.NewReader(input)
		d := Scan(r)

		if got, want := Walk(30, d), 1651; got != want {
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
		ID:       0,
		Name:     "AA",
		FlowRate: 0,
		Values:   []string{"DD", "II", "BB"},
	},
	{
		ID:       1,
		Name:     "BB",
		FlowRate: 13,
		Values:   []string{"CC", "AA"},
	},
	{
		ID:       2,
		Name:     "CC",
		FlowRate: 2,
		Values:   []string{"DD", "BB"},
	},
	{
		ID:       3,
		Name:     "DD",
		FlowRate: 20,
		Values:   []string{"CC", "AA", "EE"},
	},
	{
		ID:       4,
		Name:     "EE",
		FlowRate: 3,
		Values:   []string{"FF", "DD"},
	},
	{
		ID:       5,
		Name:     "FF",
		FlowRate: 0,
		Values:   []string{"EE", "GG"},
	},
	{
		ID:       6,
		Name:     "GG",
		FlowRate: 0,
		Values:   []string{"FF", "HH"},
	},
	{
		ID:       7,
		Name:     "HH",
		FlowRate: 22,
		Values:   []string{"GG"},
	},
	{
		ID:       8,
		Name:     "II",
		FlowRate: 0,
		Values:   []string{"AA", "JJ"},
	},
	{
		ID:       9,
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
