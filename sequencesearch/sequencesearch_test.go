package sequencesearch

import (
	"reflect"
	"testing"
)

func TestSequence(t *testing.T) {
	type args struct {
		nucleotide Nucleotide
	}
	target := "AGTA"
	prefixLen := 5
	suffixLen := 7
	n := New(target, prefixLen, suffixLen)
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				nucleotide: Nucleotide{Input: "TAGTAGGGε"},
			},
			want: []string{"T AGTA GGG"},
		},
		{
			name: "test2",
			args: args{
				nucleotide: Nucleotide{Input: "AAGTACGTGCAGTGAGTAGTAGACCTGACGTAGACCGATATAAGTAGCTAε"},
			},
			want: []string{
				"A AGTA CGTGCAG",
				"CAGTG AGTA GTAGACC",
				"TGAGT AGTA GACCTGA",
				"ATATA AGTA GCTA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n.Sequence(tt.args.nucleotide); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
