package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/marcosrachid/go-grpc/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	wd, _ := os.Getwd()
	certFile := filepath.Join(wd, "ssl", "cert.pem")
	keyFile := filepath.Join(wd, "ssl", "private.key")
	creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)

	serverAddr := fmt.Sprintf("%s:%d", store.ADDR, store.PORT)
	listen, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	store.RegisterStoreServiceService(grpcServer, store.NewServer())

	fmt.Printf("Listening gRPC on %s\n", serverAddr)
	grpcServer.Serve(listen)
}
