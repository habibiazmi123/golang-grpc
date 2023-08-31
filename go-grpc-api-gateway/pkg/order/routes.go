package order

import (
	"github.com/gin-gonic/gin"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/auth"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/config"
	"github.com/habibiazmi123/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)

	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}