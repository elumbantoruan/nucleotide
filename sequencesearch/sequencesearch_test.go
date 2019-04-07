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

func TestStringIndexFrom(t *testing.T) {
	type args struct {
		startIndex int
		source     string
		target     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test find AGTA from start index 0, where it should find the target at index 0",
			args: args{
				startIndex: 0,
				source:     "AGTAAGTA",
				target:     "AGTA"},
			want: 0,
		},
		{
			name: "test find AGTA from start index 1, where it should find the target at index 4",
			args: args{
				startIndex: 1,
				source:     "AGTAAGTA",
				target:     "AGTA"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIndexFrom(tt.args.startIndex, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("StringIndexFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
