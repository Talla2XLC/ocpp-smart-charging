package routes

import (
	"net/http"
	"ocpp-smart-charging/internal/handlers"
)

// RegisterRoutes - Routes registration
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/calculate-power-map", handlers.CalculatePowerMapHandler)
}
