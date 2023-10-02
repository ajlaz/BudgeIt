package api

func (a *API) RegisterRoutes() {
	a.POST("/insert/:userId", a.Insert)
	a.GET("/data", a.GetData)
	a.POST("/login", a.Login)
	a.POST("/register", a.Register)

}
