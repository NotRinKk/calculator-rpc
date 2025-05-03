package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type RoundingService struct{}

func (s *RoundingService) Round(args map[string]interface{}, reply *float64) error {
	value := args["value"].(float64)
	places := args["places"].(int)

	shift := math.Pow(10, float64(places))
	*reply = math.Round(value*shift) / shift
	return nil
}

func main() {
	service := new(RoundingService)
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Rounding service running on :8080")
	http.Serve(listener, nil)
}
