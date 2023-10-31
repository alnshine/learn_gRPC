package main

import (
	"context"
	"fmt"
	"grpc/adderPkg/adder"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 2 {
		log.Fatal("arg < 2")
	}
	a, err := strconv.Atoi(arg[0])
	if err != nil {
		log.Fatal("arg 1 not number")
	}
	b, err := strconv.Atoi(arg[1])
	if err != nil {
		log.Fatal("arg 2 not number")
	}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc err")
	}
	client := adder.NewAdderClient(conn)
	res, err := client.Add(context.Background(), &adder.AddRequest{X: int32(a), Y: int32(b)})
	if err != nil {
		log.Fatal("err client")
	}
	fmt.Println(res.GetRes())

}
