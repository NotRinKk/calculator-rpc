package main

import (
	"errors"
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type SquareRootService struct{}

func (s *SquareRootService) Sqrt(value float64, reply *float64) error {
	if value < 0 {
		return errors.New("square root of negative number")
	}
	*reply = math.Sqrt(value)
	return nil
}

func main() {
	service := new(SquareRootService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Square root service running on :8080")
	http.Serve(listener, nil)
}
