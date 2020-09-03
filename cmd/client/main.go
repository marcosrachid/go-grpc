package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/marcosrachid/go-grpc/internal/store"
	"github.com/marcosrachid/go-grpc/pkg/utils"
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

	serverAddr := fmt.Sprintf(
		"%s:%s",
		utils.GetenvDefault("ADDR", store.ADDR),
		utils.GetenvDefault("PORT", strconv.Itoa(store.PORT)),
	)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Fail to dial: %s\n", err)
	}

	defer conn.Close()
	client := store.NewStoreServiceClient(conn)

	ctx := context.Background()

	for {
		time.Sleep(time.Duration(3) * time.Second)

		random := rand.Intn(100) % 5
		fmt.Printf("[gRPC] random=%d\n", random)

		order := &store.OrderRequest{
			Product: store.Products(random),
		}
		response, err := client.Order(ctx, order)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		fmt.Printf("[gRPC] Received=%v\n", response.GetProduct())
	}

}
