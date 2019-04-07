package main

import (
	"context"
	"log"
	pb "nucleotide/grpc/sequencer"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	c := NewClient()
	c.Send("AAGTACGTGCAGTGAGTAGTAGACCTGACGTAGACCGATATAAGTAGCTAε")
	// c.Send("AGTAGGG")
	// c.Send("AGTAGGG")
	// c.Send("AGTAGGG")
	// c.Send("AGTAGGG")

}

// Client .
type Client struct {
}

// NewClient .
func NewClient() *Client {
	return &Client{}
}

// Send .
func (c *Client) Send(message string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewSequencerClient(conn)

	ctx := context.Background()
	stream, err := client.Next(ctx)

	nc := pb.Nucleotide{Input: message}
	err = stream.Send(&nc)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
}
