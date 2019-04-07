package sequencesearch

import (
	"fmt"
	"strings"
)

// Nucleotide .
type Nucleotide struct {
	Input string
}

//SequenceSearch .
type SequenceSearch struct {
	Target        string
	PrefixLen     int
	SuffixLen     int
	StringBuilder strings.Builder
}

// New .
func New(target string, prefixLen int, suffixLen int) *SequenceSearch {
	var sb strings.Builder
	return &SequenceSearch{
		StringBuilder: sb,
		Target:        target,
		PrefixLen:     prefixLen,
		SuffixLen:     suffixLen,
	}
}

// Sequence .
func (s *SequenceSearch) Sequence(nucleotide Nucleotide) []string {
	var (
		input      string
		results    []string
		prefix     string
		suffix     string
		output     string
		left       int
		right      int
		eofSymbol  = "ε"
		startIndex int
		eof        bool
	)

	input = nucleotide.Input

	// doesn't contain target but we need to add it into string builder
	if !strings.Contains(input, s.Target) {
		s.StringBuilder.WriteString(input)
		// after it's added then need to check if it contains the target
		if !strings.Contains(s.StringBuilder.String(), s.Target) {
			if strings.HasSuffix(input, eofSymbol) {
				// so far target never seen and it reached eof stream so clear out the builder because it's no longer needed
				s.StringBuilder.Reset()
			}
			return nil
		}
	} else {
		s.StringBuilder.WriteString(input)
	}

	// it must contain the target
	input = s.StringBuilder.String()

	// clear out string builder for efficient storage
	s.StringBuilder.Reset()

	for {
		idx := StringIndexFrom(startIndex, input, s.Target)

		if idx == -1 {
			break
		}

		left = idx - s.PrefixLen
		if idx < s.PrefixLen {
			left = 0
		}

		right = idx + len(s.Target) + s.SuffixLen
		if len(input) <= right {
			right = len(input)
		}

		prefix = input[left:idx]
		eof = strings.HasSuffix(input, eofSymbol)
		if right == len(input) && eof {
			right -= len(eofSymbol)
		}
		suffix = input[idx+len(s.Target) : right]
		output = fmt.Sprintf("%s %s %s", prefix, s.Target, suffix)
		results = append(results, output)

		startIndex = idx + len(s.Target) - 1
	}

	if !eof && right < len(input) {
		// optimization
		// put left over bytes into string builder as it may be able to use it to get next target
		// startIndex is the last index + length of target - 1
		s.StringBuilder.WriteString(input[startIndex:])
	}

	return results

}

// StringIndexFrom finds the index of substring from the startIndex instead of starts from 0.
func StringIndexFrom(startIndex int, source string, target string) int {
	sub := source[startIndex:]
	n := strings.Index(sub, target)
	if n != -1 {
		return n + startIndex
	}
	return n
}
