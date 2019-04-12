package main

import (
	"io"
	"log"
	"net"
	"nucleotide/sequencesearch"
	ss "nucleotide/sequencesearch"

	pb "nucleotide/grpc/sequencer"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := NewServer()
	pb.RegisterSequencerServer(s, srv)
	log.Println("Listening for incoming Nucleotide stream")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve a listener")
	}
}

const port = ":8080"

// Server .
type Server struct {
	Sequencer sequencesearch.SequenceSearcher
}

// NewServer creates type of server
func NewServer() *Server {
	target := "AGTA"
	prefixLen := 5
	suffixLen := 7
	seq := ss.New(target, prefixLen, suffixLen)
	return &Server{
		Sequencer: seq,
	}
}

// Next .
func (s *Server) Next(stream pb.Sequencer_NextServer) error {
	for {
		st, err := stream.Recv()
		if err == io.EOF {
			sq := pb.Sequence{
				Output: []string{},
			}
			return stream.SendAndClose(&sq)
		}
		if err != nil {
			log.Println("ERROR:")
			log.Println(err)
		}

		var result []string

		output := s.Sequencer.NextSequence(rune(st.Input))
		if len(output) > 0 {
			result = append(result, output)
		}
		s.print(result)
	}
}

func (s *Server) print(list []string) {
	if list == nil {
		return
	}
	for _, l := range list {
		log.Println(l)
	}
}
