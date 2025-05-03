package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type AdditionService struct{}

func (s *AdditionService) Add(args []float64, reply *float64) error {
	*reply = args[0] + args[1]
	return nil
}

func main() {
	service := new(AdditionService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Addition service running on :8080")
	http.Serve(listener, nil)
}
