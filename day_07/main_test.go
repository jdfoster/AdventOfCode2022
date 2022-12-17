package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	root := NewDirectory(nil)
	root.Files["b.txt"] = 14848514
	root.Files["c.txt"] = 8504156

	a := NewDirectory(root)
	a.Files["f"] = 29116
	a.Files["g"] = 2557
	a.Files["h.lst"] = 62596

	e := NewDirectory(a)
	e.Files["i"] = 584

	a.Dirs["a"] = e

	d := NewDirectory(root)
	d.Files["j"] = 4060174
	d.Files["d.log"] = 8033020
	d.Files["d.ext"] = 5626152
	d.Files["k"] = 7214296

	root.Dirs["a"] = a
	root.Dirs["d"] = d

	t.Run("Directory Sum", func(t *testing.T) {
		cases := []struct {
			input *Directory
			want  int
		}{
			{
				input: e,
				want:  584,
			},
			{
				input: a,
				want:  94853,
			},
			{
				input: d,
				want:  24933642,
			},
			{
				input: root,
				want:  48381165,
			},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
				if got := Sum(*c.input); got != c.want {
					t.Errorf("got %d, want %d", got, c.want)
				}
			})
		}
	})

	t.Run("day 6 part A", func(t *testing.T) {
		want := 95437

		if got := Walk(*root, 100_000); got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
