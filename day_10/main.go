package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var fn = "./day_10/input.txt"

type Command string

const (
	ADDX Command = "addx"
	NOOP Command = "noop"
)

type Instruction struct {
	cmd Command
	val int
}

type Register struct {
	x int
}

type Machine struct {
	cyc int
	mem *Register
}

func (m *Machine) Tick(fn func(int, Register)) {
	m.cyc++
	fn(m.cyc, *m.mem)
}

func (m *Machine) Instruct(fn func(int, Register), inst Instruction) {
	switch inst.cmd {
	case ADDX:
		m.Tick(fn)
		m.Tick(fn)
		m.mem.x += inst.val
	case NOOP:
		m.Tick(fn)
	}
}

func NewMachine() *Machine {
	return &Machine{
		mem: &Register{x: 1},
		cyc: 0,
	}
}

type Proc struct {
	bps   []int
	state map[int]Register
	crt   CRT
}

func (p Proc) Strength() int {
	var res int

	for k, v := range p.state {
		res += k * v.x
	}

	return res
}

func (p *Proc) Tick(cycle int, memory Register) {
	for _, bp := range p.bps {
		if bp != cycle {
			continue
		}

		p.state[cycle] = memory
	}

	p.crt.Draw(cycle, memory)
}

func (p *Proc) Run(mach *Machine, items []Instruction) {
	for _, item := range items {
		mach.Instruct(p.Tick, item)
	}
}

func NewProc() *Proc {
	return &Proc{
		bps:   []int{20, 60, 100, 140, 180, 220},
		state: make(map[int]Register),
		crt:   NewCRT(),
	}
}

type CRT [][]rune

func (c CRT) Draw(cycle int, mem Register) {
	x := (cycle - 1) / 40
	y := (cycle - 1) % 40

	left, right := mem.x-1, mem.x+1
	var r rune = '.'
	if y >= left && y <= right {
		r = '#'
	}

	c[x][y] = r
}

func (c CRT) Render() string {
	b := strings.Builder{}

	for _, pp := range c {
		for _, p := range pp {
			b.WriteRune(p)
		}

		b.WriteRune('\n')
	}

	return b.String()
}

func NewCRT() CRT {
	crt := make(CRT, 6)

	for i := 0; i < 6; i++ {
		crt[i] = make([]rune, 40)
	}

	return crt
}

func Scan(r io.Reader) []Instruction {
	var res []Instruction

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		switch {
		case strings.HasPrefix(line, "addx"):
			l := strings.TrimPrefix(line, "addx ")
			v, err := strconv.ParseInt(l, 10, 0)
			if err != nil {
				panic(err)
			}

			res = append(res, Instruction{cmd: ADDX, val: int(v)})
		case strings.HasPrefix(line, "noop"):
			res = append(res, Instruction{cmd: NOOP, val: 0})
		default:
			panic("failed to parse line")
		}

	}

	return res
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	items := Scan(f)
	mach := NewMachine()
	proc := NewProc()
	proc.Run(mach, items)

	fmt.Println("part 1 value: ", proc.Strength())
	fmt.Println("part 2 value:")
	fmt.Println(proc.crt.Render())
}
