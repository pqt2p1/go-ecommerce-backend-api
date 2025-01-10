package main

import "github.com/pqt2p1/go-ecommerce-backend-api/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":8080")
}
