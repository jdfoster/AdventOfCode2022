package main

import (
	"sort"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Unknown kind", func(t *testing.T) {
		cases := []struct {
			name string
			a, b Unknown
			want int
		}{
			{
				name: "should return 0 when both numbers are equal",
				want: 0,
				a:    Unknown{toNumber(1)},
				b:    Unknown{toNumber(1)},
			},
			{
				name: "should return 1 when left is less than right",
				want: 1,
				a:    Unknown{toNumber(0)},
				b:    Unknown{toNumber(1)},
			},
			{
				name: "should return -1 when left is greater than right",
				want: -1,
				a:    Unknown{toNumber(1)},
				b:    Unknown{toNumber(0)},
			},
			{
				name: "should return 0 when both list lengths are equal",
				want: 0,
				a:    Unknown{toList(toNumber(0))},
				b:    Unknown{toList(toNumber(0))},
			},
			{
				name: "should return 1 when left list is shorter than right",
				want: 1,
				a:    Unknown{toList()},
				b:    Unknown{toList(toNumber(0))},
			},
			{
				name: "should reutrn -1 when left list is longer than right",
				want: -1,
				a:    Unknown{toList(toNumber(0))},
				b:    Unknown{toList()},
			},
			{
				name: "should wrap a number in a list when the other item is a list",
				want: 0,
				a:    Unknown{toNumber(9)},
				b:    Unknown{toNumberList(9)},
			},
			{
				name: "should return expected result for example pair 1",
				want: 1,
				a:    Unknown{toNumberList(1, 1, 3, 1, 1)},
				b:    Unknown{toNumberList(1, 1, 5, 1, 1)},
			},
			{
				name: "should return expected result for example pair 2",
				want: 1,
				a:    Unknown{toList(toNumberList(1), toNumberList(2, 3, 4))},
				b:    Unknown{toList(toNumberList(1), toNumberList(4))},
			},
			{
				name: "should return expected result for example pair 3",
				want: -1,
				a:    Unknown{toNumberList(9)},
				b:    Unknown{toList(toNumberList(8, 7, 6))},
			},
			{
				name: "should return expected result for example pair 4",
				want: 1,
				a:    Unknown{toList(toNumberList(4, 4), toNumber(4), toNumber(4))},
				b:    Unknown{toList(toNumberList(4, 4), toNumber(4), toNumber(4), toNumber(4))},
			},
			{
				name: "should return expected result for example pair 5",
				want: -1,
				a:    Unknown{toNumberList(7, 7, 7, 7)},
				b:    Unknown{toNumberList(7, 7, 7)},
			},
			{
				name: "should return expected result for example pair 6",
				want: 1,
				a:    Unknown{toList()},
				b:    Unknown{toNumberList(3)},
			},
			{
				name: "should return expected result for example pair 7",
				want: -1,
				a:    Unknown{toList(toList(toList()))},
				b:    Unknown{toList(toList())},
			},
			{
				name: "should return expected result for example pair 8",
				want: -1,
				a:    Unknown{toList(toNumber(1), toList(toNumber(2), toList(toNumber(3), toList(toNumber(4), toNumberList(5, 6, 7)))), toNumber(8), toNumber(9))},
				b:    Unknown{toList(toNumber(1), toList(toNumber(2), toList(toNumber(3), toList(toNumber(4), toNumberList(5, 6, 0)))), toNumber(8), toNumber(9))},
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				if got := Compare(c.a, c.b); got != c.want {
					t.Errorf("got %d, want %d", got, c.want)
				}
			})
		}
	})

	t.Run("should parse pairs and compare", func(t *testing.T) {
		r := strings.NewReader(input)
		pp := Scan(r)

		if got, want := pp.Sum(), 13; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should sort Unknowns", func(t *testing.T) {
		r := strings.NewReader(input)
		pp := Scan(r)
		ul := pp.UnknownsList()

		lower := Unknown{toList(toNumberList(2))}
		upper := Unknown{toList(toNumberList(6))}

		ul = append(ul, lower)
		ul = append(ul, upper)

		sort.Sort(ul)

		lowerIndex, ok := ul.Find(lower)
		if !ok {
			t.Fatal("failed to find lower index")
		}
		upperIndex, ok := ul.Find(upper)
		if !ok {
			t.Fatal("failed to find upper index")
		}

		if got, want := lowerIndex*upperIndex, 140; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should calculate decode key", func(t *testing.T) {
		r := strings.NewReader(input)
		pp := Scan(r)
		ul := pp.UnknownsList()

		if got, want := DecodeKey(ul), 140; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func toNumber(v int) interface{} {
	return float64(v)
}

func toNumberList(vv ...int) interface{} {
	result := make([]interface{}, len(vv))
	for i, v := range vv {
		result[i] = toNumber(v)
	}

	return result
}

func toList(items ...interface{}) interface{} {
	result := make([]interface{}, len(items))
	for i, item := range items {
		result[i] = item
	}

	return result
}

var input = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
