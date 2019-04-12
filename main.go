package main

import (
	"log"
	ss "nucleotide/sequencesearch"
)

func main() {

	target := "AGTA"
	x := 5
	y := 7

	s := "GGGGGGGGGGGGGGGGGGGAAGTACGTGCAGTGAGTAGTAGACCTGACGTAGACCGATATAAGTAGCTAε"

	runes := []rune(s)
	nc := ss.New(target, x, y)
	for i := 0; i < len(runes); i++ {
		res := nc.NextSequence(runes[i])
		if len(res) > 0 {
			log.Println(string(res))
		}
	}
}
