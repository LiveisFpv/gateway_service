package grpc

import "google.golang.org/grpc"

//TODO connection to GRPC server

type Client struct {
	port string
	app  *grpc.ClientConn
}
