package api

func (a *API) RegisterRoutes() {
	a.POST("/transactions", a.Insert)
	a.GET("/transactions/:userId", a.GetData)
	a.GET("/transactions/:userId/:num", a.GetNum)
	a.POST("/login", a.Login)
	a.POST("/register", a.Register)


}
