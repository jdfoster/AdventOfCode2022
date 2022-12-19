package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func Sum(dir *Directory) int {
	var f func(*Directory) int

	f = func(d *Directory) int {
		var acc int

		for _, f := range d.Files {
			acc += f
		}

		for _, d := range d.Dirs {
			if d == nil {
				continue
			}

			acc += f(d)
		}

		return acc
	}

	return f(dir)
}

func Walk(dir *Directory, cutoff int) (int, Directories) {
	var (
		walk  func(*Directory)
		total int
		dirs  []*Directory
	)

	walk = func(d *Directory) {
		for _, d := range d.Dirs {
			if d == nil {
				continue
			}
			walk(d)

			dirs = append(dirs, d)

			if v := Sum(d); v <= cutoff {
				total += v
			}
		}
	}

	walk(dir)

	return total, dirs
}

type Directories []*Directory

func (dd Directories) Len() int {
	return len(dd)
}

func (dd Directories) Swap(i, j int) {
	dd[i], dd[j] = dd[j], dd[i]
}

func (dd Directories) Less(i, j int) bool {
	return Sum(dd[i]) < Sum(dd[j])
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

	first, dirs := Walk(root, 100_000)
	fmt.Println("part one value: ", first)

	sort.Sort(dirs)
	free := 70000000 - Sum(root)
	target := 30000000 - free

	for _, d := range dirs {
		if v := Sum(d); v >= target {
			fmt.Println("part two value: ", v)
			break
		}
	}
}
