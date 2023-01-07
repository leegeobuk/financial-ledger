package api

import (
	"github.com/leegeobuk/financial-ledger/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *API) setRoutes() {
	docs.SwaggerInfo.Title = "Financial Ledger"
	docs.SwaggerInfo.Description = "API document for Financial Ledger"
	docs.SwaggerInfo.Version = "1.0"
	ledger := a.router.Group("/api/ledger")

	// API Document (swagger)
	ledger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ledger.GET("/ping", a.Ping)
	ledger.GET("/time", a.Time)
}
