package main

import "net/http"

func (a *api) routes() http.Handler {

	router := http.NewServeMux()

	router.Handle("POST /api/newService", a.handlerCreateService())

	return router
}
