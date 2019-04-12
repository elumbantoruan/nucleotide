package sequencesearch

import (
	"fmt"
	"strings"
)

// SequenceSearcher is an interface to search a next sequence
type SequenceSearcher interface {
	NextSequence(rune) string
}

// SequenceSearch is a concrete implementation of SequenceSearcher which searchs a sequence in the incoming nucleotide by finding the target,
// and add prefix and suffix (if applicable) from a given  prefix length and suffix length respectively.
// StringBuilder is used as a buffer for an incoming input, so it can maintain the continuous stream as it's building a sequence.
// Note there are some optimization with the buffer to minimize its footprint.
// The buffer will be truncated once it reaches the end of stream and after it's being read into a string.
// At the end, the buffer will be populated with some left over input data that has been used as a target, but potentially
// it can be used for the next incoming stream
type SequenceSearch struct {
	Target          string
	Prefix          string
	Suffix          string
	PrefixTargetLen int
	SuffixTargetLen int
	StringBuilder   strings.Builder
	FoundTarget     bool
	PrefixTargeted  bool
	SuffixTargeted  bool
	EOF             bool
	TargetIndex     int
}

// New creates an instance of SequenceSearcher
func New(target string, prefixTargetLen int, suffixTargetLen int) SequenceSearcher {
	var sb strings.Builder
	return &SequenceSearch{
		StringBuilder:   sb,
		Target:          target,
		PrefixTargetLen: prefixTargetLen,
		SuffixTargetLen: suffixTargetLen,
	}
}

// NextSequence builds the sequence
func (s *SequenceSearch) NextSequence(input rune) string {
	var (
		eofSymbol = 'ε'

		bufferLength int
		buffer       string
		output       string
	)

	if input == eofSymbol {
		s.EOF = true
	} else {
		s.StringBuilder.WriteRune(input)
	}
	bufferLength = s.StringBuilder.Len()
	buffer = s.StringBuilder.String()

	if !s.FoundTarget {
		s.FoundTarget = s.isTargetFound(buffer)
	}

	// remove buffer 1 character from the front since it hasn't found the target but bufferLength has exceeded prefix targeted length + len(target)
	if !s.PrefixTargeted && !s.FoundTarget && bufferLength > s.PrefixTargetLen+len(s.Target) {
		buffer = buffer[1:bufferLength]
		s.StringBuilder.Reset()
		s.StringBuilder.WriteString(buffer)
		return ""
	}

	// target is found, next is to assign the prefix
	if s.FoundTarget && !s.PrefixTargeted {
		s.TargetIndex = strings.Index(buffer, s.Target)
		s.Prefix = buffer[:s.TargetIndex]
		if len(s.Prefix) > s.PrefixTargetLen {
			begin := len(s.Prefix) - s.PrefixTargetLen
			s.Prefix = s.Prefix[begin:s.TargetIndex]
			buffer = buffer[begin:]
			s.StringBuilder.Reset()
			s.StringBuilder.WriteString(buffer)
		}

		// prefix targeted can be empty string, less than or match the prefix targeted length
		s.PrefixTargeted = true

		// target is found and prefix has been allocated, next is to assign the suffix
	} else if s.FoundTarget && s.PrefixTargeted {
		maxLength := len(s.Prefix) + len(s.Target) + s.SuffixTargetLen

		if (bufferLength == maxLength) || s.EOF {
			if s.EOF {
				s.Suffix = buffer[len(s.Prefix)+len(s.Target) : bufferLength]
			} else {
				s.Suffix = buffer[len(s.Prefix)+len(s.Target) : maxLength]
			}
			s.SuffixTargeted = true
			output = fmt.Sprintf("%s %s %s", s.Prefix, s.Target, s.Suffix)

			// truncate since it's the end of stream
			if s.EOF {
				s.reset()
				s.StringBuilder.Reset()
			} else {
				// need to find other target in the buffer (prefix + target + suffix)
				// for overlapping targets
				idx := StringIndexFrom(s.TargetIndex+1, buffer, s.Target)
				if idx >= 0 {
					left := idx - s.PrefixTargetLen
					if left < 0 {
						left = 0
					}
					buffer = buffer[left:]
					s.StringBuilder.Reset()
					s.StringBuilder.WriteString(buffer)

					// find the next prefix and target
					s.PrefixTargeted = true
					s.FoundTarget = true
					if idx > s.PrefixTargetLen {
						s.Prefix = buffer[:s.PrefixTargetLen]
					} else {
						s.Prefix = buffer[:idx]
					}
					s.Suffix = ""

				} else {

					if len(buffer) > s.PrefixTargetLen {
						buffer = buffer[len(buffer)-s.PrefixTargetLen:]
					}

					s.StringBuilder.Reset()
					s.StringBuilder.WriteString(buffer)

					s.PrefixTargeted = false
					s.FoundTarget = false
				}
			}

		}

	}

	return output
}

func (s *SequenceSearch) isTargetFound(input string) bool {
	return strings.Contains(input, s.Target)
}

func (s *SequenceSearch) reset() {
	s.Prefix = ""
	s.Suffix = ""
	s.FoundTarget = false
	s.PrefixTargeted = false
	s.SuffixTargeted = false
	s.EOF = false
}

// StringIndexFrom finds the index of substring from the startIndex instead of starts from 0 (the builtin strings.Index)
func StringIndexFrom(startIndex int, source string, target string) int {
	sub := source[startIndex:]
	n := strings.Index(sub, target)
	if n != -1 {
		return n + startIndex
	}
	return n
}
