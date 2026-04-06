package weather

import (
	"net/http"
	"weather-api/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("GET /weather", h.GetWeather)
}
