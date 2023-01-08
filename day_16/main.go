package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fn = "./day_16/input.txt"

type Walker struct {
	Current *Valve
	VState  uint
}

func (w Walker) Copy() *Walker {
	return &Walker{
		Current: w.Current,
		VState:  w.VState,
	}
}

func (w Walker) Inspect() (int, bool) {
	var mask uint = 1 << w.Current.ID
	open := (w.VState & mask) == 0
	return w.Current.FlowRate, open
}

func (w *Walker) Open() {
	var mask uint = 1 << w.Current.ID
	w.VState = w.VState | mask
}

func (w *Walker) Move(name string) bool {
	for i := 0; i < len(w.Current.Tunnel); i++ {
		nxt := w.Current.Tunnel[i]
		if nxt.Name == name {
			w.Current = nxt
			return true
		}
	}

	return false
}

type Records map[Walker]int

func Walk(graph *Valve, dur int) int {
	steps := make([]Records, dur+1)
	for i := 0; i < len(steps); i++ {
		steps[i] = make(Records)
	}

	put := func(t int, w Walker, p int) {
		if cur, ok := steps[t][w]; !ok || p > cur {
			steps[t][w] = p
		}
	}

	put(0, Walker{Current: graph}, 0)

	for i := 0; i < dur; i++ {
		step := steps[i]
		fmt.Printf("time: %02d; count: %d\n", i, len(step))

		for k, v := range step {
			if r, ok := k.Inspect(); ok {
				k.Open()
				auc := (dur - i - 1) * r
				put(i+1, k, v+auc)
			}

			for _, t := range k.Current.Tunnel {
				w := k.Copy()
				if w.Move(t.Name) {
					put(i+1, *w, v)
				}
			}
		}
	}

	last := steps[dur]
	var max int = math.MinInt

	for _, v := range last {
		if v > max {
			max = v
		}
	}

	return max
}

type Valve struct {
	ID       uint
	Name     string
	Tunnel   []*Valve
	FlowRate int
}

type ValueDetail struct {
	ID       uint
	Name     string
	FlowRate int
	Values   []string
}

type ValueDetails map[string]ValueDetail

func (dd ValueDetails) Build(root string) *Valve {
	pts := make(map[string]*Valve, len(dd))

	for k, v := range dd {
		pts[k] = &Valve{
			ID:       v.ID,
			Name:     v.Name,
			FlowRate: v.FlowRate,
		}
	}

	link := func(parent string, children []string) {
		tunnels := make([]*Valve, len(children))

		for i, child := range children {
			v, ok := pts[child]
			if !ok {
				panic("failed to find child valve :" + child)
			}

			tunnels[i] = v
		}

		p, ok := pts[parent]
		if !ok {
			panic("failed to find parent value :" + parent)
		}

		p.Tunnel = tunnels
	}

	for k, v := range dd {
		link(k, v.Values)
	}

	r, ok := pts[root]
	if !ok {
		panic("failed to find root value :" + root)
	}

	return r
}

func Scan(r io.Reader) ValueDetails {
	re := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
	s := bufio.NewScanner(r)

	results := make(map[string]ValueDetail)

	var id uint
	for s.Scan() {
		mm := re.FindAllStringSubmatch(s.Text(), -1)
		if len(mm) > 0 && len(mm[0]) != 4 {
			panic("failed to parse line")
		}

		var (
			err  error
			meta ValueDetail
		)

		meta.ID = id
		meta.Name = mm[0][1]

		meta.FlowRate, err = strconv.Atoi(mm[0][2])
		if err != nil {
			panic("failed to parse number")
		}

		meta.Values = strings.Split(mm[0][3], ",")

		for i := 1; i < len(meta.Values); i++ {
			meta.Values[i] = strings.TrimSpace(meta.Values[i])
		}

		results[meta.Name] = meta

		id++
	}

	return results
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	d := Scan(f)
	g := d.Build("AA")
	first := Walk(g, 30)

	fmt.Println("part 1 value: ", first)
}
