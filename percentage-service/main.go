package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type PercentageService struct{}

func (s *PercentageService) Percentage(args map[string]float64, reply *float64) error {
	*reply = args["value"] * args["percent"] / 100
	return nil
}

func main() {
	service := new(PercentageService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Percentage service running on :8080")
	http.Serve(listener, nil)
}
