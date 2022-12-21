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
					validateMonkey(t, barrel[i], tc)
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

		t.Run("should iterate a round for 20 iterations", func(t *testing.T) {
			r := strings.NewReader(input)
			barrel := Scan(r)

			for i := 0; i < 20; i++ {
				barrel.Round(3)
			}

			if got, want := barrel.Business(), 10605; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})

		t.Run("should iterate a round for 1000 iterations", func(t *testing.T) {
			r := strings.NewReader(input)
			barrel := Scan(r)

			count := 0

			runTo := func(limit int, counts []int) {
				if count > limit {
					t.Fatal("counter has passed limit")
				}

				for count < limit {
					barrel.Round(1)
					count++
				}

				t.Logf("Round %d\n", count)
				validateCounts(t, barrel, counts)
			}

			runTo(1, []int{2, 4, 3, 6})
			runTo(20, []int{99, 97, 8, 103})
			runTo(1000, []int{5204, 4792, 199, 5192})
			runTo(2000, []int{10419, 9577, 392, 10391})
			runTo(3000, []int{15638, 14358, 587, 15593})
			runTo(4000, []int{20858, 19138, 780, 20797})
			runTo(5000, []int{26075, 23921, 974, 26000})
			runTo(6000, []int{31294, 28702, 1165, 31204})
			runTo(7000, []int{36508, 33488, 1360, 36400})
			runTo(8000, []int{41728, 38268, 1553, 41606})
			runTo(9000, []int{46945, 43051, 1746, 46807})
			runTo(10000, []int{52166, 47830, 1938, 52013})

			if got, want := barrel.Business(), 2713310158; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	})
}

func validateCounts(t *testing.T, b Barrel, counts []int) {
	t.Helper()

	if got, want := len(b), len(counts); got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	for i, want := range counts {
		mky := b[i]
		if got := mky.count; got != want {
			t.Fatalf("monkey %d; got %d, want %d", i, got, want)
		}
	}
}

func validateMonkey(t *testing.T, mky *Monkey, tc Monkey) {
	t.Helper()

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
