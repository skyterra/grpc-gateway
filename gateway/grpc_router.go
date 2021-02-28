package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
	"time"

	pb "xframework/pb"
)

const (
	defaultBodySize = 1 * 1024 * 1024
)

func GrpcRouter(w http.ResponseWriter, r *http.Request) {
	param := make([]byte, defaultBodySize)
	r.Body.Read(param)

	fields := strings.Split(r.URL.Path, "/")
	servName := fields[1]

	req := &pb.ServRequest{
		Operation: fields[2],
		Method:    r.Method,
		Query:     r.URL.RawQuery,
		Body:      string(param),
	}

	conn, err := grpc.Dial(servName, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroServClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Handle(ctx, req)
	if err != nil {
		log.Fatalf("could not finish job: %v", err)
	}

	w.Write([]byte(resp.GetResponse()))
}
