package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	pb "nucleotide/grpc/sequencer"
	"os"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	path := flag.String("path", "input.txt", "path for input file")
	flag.Parse()

	var (
		input []string
		err   error
	)

	c := NewClient()
	if *path != "" {
		input, err = readLines(*path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if len(os.Args) == 2 {
			input = []string{os.Args[1]}
		} else {
			log.Fatal("enter argument")
		}
	}
	for _, line := range input {
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			c.Send(runes[i])
		}
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Client .
type Client struct {
}

// NewClient .
func NewClient() *Client {
	return &Client{}
}

// Send .
func (c *Client) Send(char int32) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewSequencerClient(conn)

	ctx := context.Background()
	stream, err := client.Next(ctx)
	if err != nil {
		log.Println(err)
		log.Println("Make sure if the server is up and running")
		return
	}

	nc := pb.Nucleotide{Input: char}
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
