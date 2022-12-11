package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("day 6 part A", func(t *testing.T) {
		cases := []struct {
			input string
			want  int
		}{
			{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				want:  7,
			},
				{
					input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
					want:  5,
				},
				{
					input: "nppdvjthqldpwncqszvftbrmjlhg",
					want:  6,
				},
				{
					input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
					want:  10,
				},
				{
					input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
					want:  11,
				},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
				if got := findPacket([]byte(c.input)); got != c.want {
					t.Errorf("got %d, want %d", got, c.want)
				}
			})
		}
	})
}
