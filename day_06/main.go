package main

import (
	"fmt"
	"os"
)

var fn = "./day_06/input.txt"

const (
	packetStartSize = 4
	messageStartSize = 14
)

func toByteSet(bb []byte) map[byte]struct{} {
	res := make(map[byte]struct{})

	for _, b := range bb {
		res[b] = struct{}{}
	}

	return res
}

func findPacket(l int, b []byte) int {
	for j := l; j < len(b); j++ {
		i := j - l
		s := toByteSet(b[i:j])

		if len(s) == l {
			return j
		}
	}

	return -1
}

func main() {
	b, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	first := findPacket(packetStartSize, b)
	second := findPacket(messageStartSize, b)
	fmt.Println("part one value: ", first)
	fmt.Println("part two value: ", second)
}
