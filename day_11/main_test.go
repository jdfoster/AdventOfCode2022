package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Scan", func(t *testing.T) {
		t.Run("should parse input", func(t *testing.T) {
			cases := []Monkey{
				{
					items: []int{79, 98},
					div:   23,
					pass:  2,
					fail:  3,
					op:    buildMulti(19),
				},
				{
					items: []int{54, 65, 75, 74},
					div:   19,
					pass:  2,
					fail:  0,
					op:    buildAdd(6),
				},
				{
					items: []int{79, 60, 97},
					div:   13,
					pass:  1,
					fail:  3,
					op:    buildSquare(),
				},
				{
					items: []int{74},
					div:   17,
					pass:  0,
					fail:  1,
					op:    buildAdd(3),
				},
			}

			r := strings.NewReader(input)
			barrel := Scan(r)

			if got, want := len(barrel), 4; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}

			for i, tc := range cases {
				t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
					ValidateMonkey(t, barrel[i], tc)
				})
			}
		})
	})

	t.Run("Barrel", func(t *testing.T) {
		t.Run("should calculate monkey business", func(t *testing.T) {
			res := Barrel{{count: 101}, {count: 95}, {count: 7}, {count: 105}}
			if got, want := res.Business(), 10605; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})

		t.Run("should iterate a round", func(t *testing.T) {
			r := strings.NewReader(input)
			barrel := Scan(r)

			for i := 0; i < 20; i++ {
				barrel.Round()
			}

			if got, want := barrel.Business(), 10605; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	})
}

func ValidateMonkey(t *testing.T, mky *Monkey, tc Monkey) {
	if got, want := len(mky.items), len(tc.items); got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	for i, got := range mky.items {
		if want := tc.items[i]; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	if got, want := mky.div, tc.div; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	if got, want := mky.pass, tc.pass; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	if got, want := mky.fail, tc.fail; got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	for i := 0; i < 10; i++ {
		if got, want := mky.op(i), tc.op(i); got != want {
			t.Errorf("insect; got %d, want %d", got, want)
		}
	}
}

var input = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
