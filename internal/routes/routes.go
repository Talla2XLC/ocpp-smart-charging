package routes

import (
	"net/http"
	"ocpp-smart-charging/internal/handlers"
)

// RegisterRoutes - Routes registration
func RegisterRoutes() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/calculate-power-map", handlers.CalculatePowerMapHandler)
}
