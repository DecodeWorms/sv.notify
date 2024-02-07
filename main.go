package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"sv.notify/config"
	"sv.notify/constant"
	"sv.notify/grpc"
)

func main() {
	cfg := config.ImportConfig(config.Config{})
	grpcServ, err := grpc.NewServer(constant.WelcomeTopic, cfg.PulsarUrl, cfg.PulsarServiceName)
	if err != nil {
		log.Panic(err)
	}

	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	go func(addr string) {
		log.Printf("Notify gRPC service running on %v. Environment=%s", addr, "")
		if err := grpcServ.Run(addr); err != nil {
			log.Println("failed to start server")
		}
	}(addr)

	<-interruptHandler
	log.Println("closing application...")

}
