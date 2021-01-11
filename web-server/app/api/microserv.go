package api

import (
	"context"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "xframework/protobuf"
)

var ServMap map[string]string = map[string]string{
	"serv-1": "localhost:50051",
	"serv-2": "localhost:50052",
}

type MicroServRequest struct {
	ServName string
	Params   string
}

func MicroServDo(r *ghttp.Request) {
	var data *MicroServRequest
	if err := r.Parse(&data); err != nil {
		log.Fatal(err.Error())
	}

	address, exist := ServMap[data.ServName]
	if !exist {
		log.Fatalf("micro-server %s not exist", data.ServName)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMicroServClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.DoJob(ctx, &pb.ServRequest{Label: data.ServName, Params: data.Params})
	if err != nil {
		log.Fatalf("could not finish job: %v", err)
	}
	r.Response.Writeln(fmt.Sprintf("MicroServ:%s", resp.GetResponse()))
}
