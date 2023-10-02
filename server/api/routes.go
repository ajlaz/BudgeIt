package api

func (a *API) RegisterRoutes() {
	a.POST("/insert", a.Insert)
	a.GET("/data/:userId", a.GetData)
	a.POST("/login", a.Login)
	a.POST("/register", a.Register)

}
