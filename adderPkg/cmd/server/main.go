package main

import (
	"grpc/adderPkg/adder"
	"grpc/adderPkg/pkg"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &pkg.GRPCServer{}
	adder.RegisterAdderServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
