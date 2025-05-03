package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type SubtractionService struct{}

func (s *SubtractionService) Subtract(args []float64, reply *float64) error {
	*reply = args[0] - args[1]
	return nil
}

func main() {
	service := new(SubtractionService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Subtraction service running on :8080")
	http.Serve(listener, nil)
}
