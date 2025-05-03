package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type DivisionService struct{}

func (s *DivisionService) Divide(args []float64, reply *float64) error {
	if args[1] == 0 {
		return errors.New("division by zero")
	}
	*reply = args[0] / args[1]
	return nil
}

func main() {
	service := new(DivisionService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Division service running on :8080")
	http.Serve(listener, nil)
}
