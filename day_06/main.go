package main

import (
	"fmt"
	"os"
)

var fn = "./day_06/input.txt"

const packetSize = 4

func toByteSet(bb []byte) map[byte]struct{} {
	res := make(map[byte]struct{})

	for _, b := range bb {
		res[b] = struct{}{}
	}

	return res
}

func findPacket(b []byte) int {
	for j := packetSize; j < len(b); j++ {
		i := j - packetSize
		s := toByteSet(b[i:j])

		if len(s) == packetSize {
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

	first := findPacket(b)
	fmt.Println("part one value: ", first)
}
