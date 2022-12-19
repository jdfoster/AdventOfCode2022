package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fn = "./day_07/input.txt"

const (
	cd  = "$ cd "
	ls  = "$ ls"
	dir = "dir "
)

type Directory struct {
	Parent *Directory
	Dirs   map[string]*Directory
	Files  map[string]int
}

func NewDirectory(p *Directory) *Directory {
	return &Directory{
		Parent: p,
		Dirs:   make(map[string]*Directory),
		Files:  make(map[string]int),
	}
}

func Sum(dir Directory) int {
	var f func(Directory) int

	f = func(d Directory) int {
		var acc int

		for _, f := range d.Files {
			acc += f
		}

		for _, d := range d.Dirs {
			if d == nil {
				continue
			}

			acc += f(*d)
		}

		return acc
	}

	return f(dir)
}

func Walk(dir Directory, cutoff int) int {
	var (
		walk  func(Directory)
		total int
	)

	walk = func(d Directory) {
		for _, d := range d.Dirs {
			if d == nil {
				continue
			}
			walk(*d)

			if v := Sum(*d); v <= cutoff {
				total += v
			}
		}
	}

	walk(dir)

	return total
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	root := NewDirectory(nil)
	var cur *Directory

	m := regexp.MustCompile(`^(\d+) (.*)`)

	for s.Scan() {
		line := s.Text()

		switch {
		case strings.HasPrefix(line, cd):
			d := strings.TrimPrefix(line, cd)
			if d == "/" {
				cur = root
				break
			}

			if d == ".." {
				cur = (*cur).Parent
				break
			}

			dd, ok := (*cur).Dirs[d]
			if !ok {
				dd = NewDirectory(cur)
			}

			(*cur).Dirs[d] = dd
			cur = dd

		case strings.HasPrefix(line, ls):

		case strings.HasPrefix(line, dir):
			d := strings.TrimPrefix(line, dir)
			if _, ok := (*cur).Dirs[d]; !ok {
				(*cur).Dirs[d] = NewDirectory(cur)
			}

		case m.MatchString(line):
			v := m.FindAllStringSubmatch(line, -1)
			if len(v) != 1 && len(v[0]) != 2 {
				panic("failed to parse with regex")
			}

			s, f := v[0][1], v[0][2]
			ss, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				panic("failed to parse int")
			}

			(*cur).Files[f] = int(ss)
		default:
			panic(fmt.Sprintf("unmatched line: %s", line))
		}
	}

	var first int
	for _, d := range root.Dirs {
		first += Walk(*d, 100_000)
	}

	fmt.Println("part one value: ", first)
}
