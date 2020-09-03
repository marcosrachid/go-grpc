package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/grpc/peer"
)

const (
	ADDR string = "localhost"
	PORT int    = 4000
)

var PRODUCTS = map[Products]Product{
	Products_SHIRT: Product{
		Product: Products_SHIRT,
		Value:   19.99,
	},
	Products_PANTS: Product{
		Product: Products_PANTS,
		Value:   99.99,
	},
	Products_SOCKS: Product{
		Product: Products_SOCKS,
		Value:   1.99,
	},
	Products_UNDERWEAR: Product{
		Product: Products_UNDERWEAR,
		Value:   4.99,
	},
	Products_DRESS: Product{
		Product: Products_DRESS,
		Value:   199.99,
	},
	Products_SHOES: Product{
		Product: Products_SHOES,
		Value:   79.99,
	},
}

type StoreServer struct {
}

func NewServer() *StoreServiceService {
	rand.Seed(time.Now().UnixNano())

	server := &StoreServer{}
	return NewStoreServiceService(server)
}

func (s *StoreServer) Order(ctx context.Context, r *OrderRequest) (*OrderResponse, error) {
	p, _ := peer.FromContext(ctx)
	fmt.Printf("[gRPC] [$=%s] Order=%s\n", p.Addr.String(), r.GetProduct())

	product := PRODUCTS[r.GetProduct()]
	response := &OrderResponse{
		Product:      &product,
		DeliveryDate: time.Now().UnixNano(),
	}
	return response, nil
}
