package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/auth"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/config"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/order"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
