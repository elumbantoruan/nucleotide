# nucleotide

## Purpose

The intent of this project is to build a subsequences of [nucleotide](https://en.m.wikipedia.org/wiki/Nucleotide).  
The input comes as a stream, which takes on one of the values: A, C, G, T, or special value ε to indicate as end of stream.  
The output is a list of sequences that are made by prefix, target, and suffix.  

## Dependency

This project requires go version >= 1.11 as it utilizes Go Modules for the dependency management.

## Project structure

### sequencesearch

#### type SequenceSearch

SequenceSearch search a sequence in the incoming nucleotide by finding the target, and add prefix and suffix (if applicable)
from a given  prefix length and suffix length respectively.
StringBuilder is used as a buffer for an incoming input, so it can maintain the continuous stream as it's building a sequence.
Note there are some optimization with the buffer to minimize its footprint.
The buffer will be truncated once it reaches the end of stream and after it's being read into a string.
At the end, the buffer will be populated with some left over input data that has been used as a target, but potentially
it can be used for the next incoming stream

``` go
type SequenceSearcher interface {
    NextSequence(rune) string
}

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
```

#### func New()

New creates an instance of SequenceSearch

``` go
func New(target string, prefixLen int, suffixLen int) *SequenceSearch
```

#### func (*SequenceSearch) NextSequence

Sequence builds a sequence of nucleotide

``` go
func (s *SequenceSearch) NextSequence(input rune) string
```

#### func StringIndexFrom

StringIndexFrom searches the target in the source from a given index.  This function is needed because the builtin strings.Index only searches from index 0, while it's necessary to continuously find other targets in the source.

``` go
func StringIndexFrom(startIndex int, source string, target string) int
```

### grpc

This package is the implementation of streaming of nucleotide in gRPC.  It contains the client, sequencer, and server sub packages.  The client sends the stream in protobuf format with the implementation of client interface generated for gRPC.  The sequencer contains a generated code of protobuf structure along with the interfaces for client and server.  The server receives the stream by implementing a server interface generated for gRPC and it utilizes the sequencesearch package to generate sequences.

#### client

The input data can be set from a command line argument or simply put a path for an input file.  
From gRPC/client folder, build the client `go build` or simply run `go run main.go`

``` go
   ./client AGTAAGTA
```

or

``` go
   ./client -path=input.txt
```

Make sure the Server is running before executing the client.

#### server

From gRPC/server folder, build the server `go build` or simply run `go run main.go`
For simplicity, the server just prints the sequences into a console.  It doesn't send the stream back to the client.
The reason is because the EOF of stream is controlled by a special character located in the stream itself.

### main

main.go contains the simple command to execute sequencesearch functionality without having the need to run client/server gRPC components.
