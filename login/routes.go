package main

func (a *application) registerRoutes() {
	api := a.server.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/login", a.loginHandler)
}
