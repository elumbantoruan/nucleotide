package main

import (
	"log"
	ss "nucleotide/sequencesearch"
)

func main() {
	source := []string{
		"A",
		"AAGTACGTGCAGTGAGTAGTAGACCTGACGTAGACCGATATAAGTAGCTAGGGA",
		"GTA",
	}
	target := "AGTA"
	x := 5
	y := 7
	var sequences []string
	var nc = ss.New(target, x, y)
	for i := 0; i < len(source); i++ {
		res := nc.NextSequence(ss.Nucleotide{Input: source[i]})
		sequences = append(sequences, res...)
	}
	log.Println(sequences)
}
