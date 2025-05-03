package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type PowerService struct{}

func (s *PowerService) Power(args map[string]float64, reply *float64) error {
	*reply = math.Pow(args["base"], args["exponent"])
	return nil
}

func main() {
	service := new(PowerService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Power service running on :8080")
	http.Serve(listener, nil)
}
