package main

import (
	"log"
	"nucleotide/sequencesearch"
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
	var nc = sequencesearch.New(target, x, y)
	for i := 0; i < len(source); i++ {
		res := nc.Sequence(ss.Nucleotide{Input: source[i]})
		sequences = append(sequences, res...)
	}
	log.Println(sequences)
}
