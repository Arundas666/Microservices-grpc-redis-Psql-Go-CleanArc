package api

import (
	"auth_svc/pkg/config"
	"auth_svc/pkg/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewGrpcServer(c *config.Config, service pb.AuthServiceServer) (*Server, error) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Gs:   grpcServer,
		Lis:  lis,
		Port: c.Port,
	}, nil

}

func (s *Server) Start() error{

	fmt.Println("Authentication service on:", s.Port)
	return s.Gs.Serve(s.Lis)

}
