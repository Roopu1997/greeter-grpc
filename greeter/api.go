package greeter

import (
	context "context"
	"log"
)

// Server greeter server
type Server struct{}

// Greet - Greet implementation
func (s *Server) Greet(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	log.Println("Request from " + req.Name)
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}
