package api

import (
	"github.com/leegeobuk/household-ledger/api/mw"
	"github.com/leegeobuk/household-ledger/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) setRoutes() {
	docs.SwaggerInfo.Title = "Household Ledger"
	docs.SwaggerInfo.Description = "API document for Household Ledger"
	docs.SwaggerInfo.Version = "v1.0"
	ledger := s.router.Group("/api/household-ledger")

	// API Document (swagger)
	ledger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ledger.GET("/ping", s.Ping)
	ledger.GET("/time", s.Time)

	// user
	ledger.POST("/user/signup", mw.ValidateSignUp, s.SignUp)
	ledger.POST("/user/signin", mw.ValidateSignIn, s.SignIn)

	// ledger
	ledger.GET("/ledgers/:user_id", mw.ValidateGetLedgers, s.GetLedgers)
	ledger.GET("/ledger/:ledger_id", mw.ValidateGetLedger, s.GetLedger)
	ledger.POST("/ledger", mw.ValidateAddLedger, s.AddLedger)
}
