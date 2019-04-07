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
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test where no sequence is found",
			args: args{
				nucleotide: Nucleotide{Input: "AAAAAAAA"},
			},
			want: nil,
		},
		{
			name: "test where target is found along with prefix and suffix generated in one line",
			args: args{
				nucleotide: Nucleotide{Input: "TAGTAGGGε"},
			},
			want: []string{"T AGTA GGG"},
		},
		{
			name: "test where target are found along with prefix and suffix generated in multiple lines",
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
			// each test data runs in a new instance
			n := New(target, prefixLen, suffixLen)
			if got := n.NextSequence(tt.args.nucleotide); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
