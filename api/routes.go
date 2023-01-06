package api

func (a *API) setRoutes() {
	ledger := a.router.Group("/api/ledger")

	ledger.GET("/echo", a.Echo)
	ledger.GET("/time", a.Time)
}
