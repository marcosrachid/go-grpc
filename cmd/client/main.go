package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/marcosrachid/go-grpc/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	wd, _ := os.Getwd()
	certFile := filepath.Join(wd, "ssl", "cert.pem")
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("Error creating credentials: %s\n", err)
	}

	serverAddr := fmt.Sprintf("%s:%d", store.ADDR, store.PORT)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Fail to dial: %s\n", err)
	}

	defer conn.Close()
	client := store.NewStoreServiceClient(conn)

	ctx := context.Background()

	for {
		order := &store.OrderRequest{
			Product: store.Products(rand.Intn(100) % 5),
		}
		response, err := client.Order(ctx, order)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		fmt.Printf("[gRPC] Received=%s\n", response.GetProduct())
	}

}
