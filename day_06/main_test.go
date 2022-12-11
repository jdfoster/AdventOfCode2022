package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	cases := []struct {
		input  string
		first  int
		second int
	}{
		{
			input:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			first:  7,
			second: 19,
		},
		{
			input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			first:  5,
			second: 23,
		},
		{
			input:  "nppdvjthqldpwncqszvftbrmjlhg",
			first:  6,
			second: 23,
		},
		{
			input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			first:  10,
			second: 29,
		},
		{
			input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			first:  11,
			second: 26,
		},
	}

	t.Run("day 6 part A", func(t *testing.T) {
		for i, c := range cases {
			t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
				if got, want := findPacket(packetStartSize, []byte(c.input)), c.first; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			})
		}
	})

	t.Run("day 6 part B", func(t *testing.T) {
		for i, c := range cases {
			t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
				if got, want := findPacket(messageStartSize, []byte(c.input)), c.second; got != want {
					t.Errorf("got %d, want %d", got, want)
				}
			})
		}
	})
}
