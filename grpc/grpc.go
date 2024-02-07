package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"sv.notify/config"
	"sv.notify/event"
	"sv.notify/pb/protos/pb/notify"
)

type Server struct {
	grpc *grpc.Server
	notify.UnimplementedNotifyServiceServer
}

func NewServer(topic, PulsarUrl, serviceName string) (*Server, error) {
	grpc := grpc.NewServer()
	ev, err := event.NewEvents(PulsarUrl, serviceName)
	if err != nil {
		return nil, err
	}
	go func() {
		err := ev.SendWelcomeMessage(topic)
		if err != nil {
			log.Printf("Error sending welcome message: %v", err)
			// Handle the error as needed
		}
	}()
	return &Server{
		grpc: grpc,
	}, nil
}

func (s *Server) Run(add string) error {
	cfg := config.ImportConfig(config.Config{})
	list, err := net.Listen(cfg.ServerProtocol, add)
	if err != nil {
		return err
	}
	return s.grpc.Serve(list)
}
