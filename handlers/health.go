package handlers

import (
	"github.com/go-chi/render"
	"net/http"
)

func (r *Router) GetHealthStatus(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"ping": "pong"})
}
