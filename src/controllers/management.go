package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ManagementController handles management related endpoints.
type ManagementController struct {
}

// HealthStatus represents the health status of a component.
type HealthStatus struct {
	Status string `json:"status"`
}

// HealthResponse represents the response format for health endpoints.
type HealthResponse struct {
	Status     string                  `json:"status"`
	Components map[string]HealthStatus `json:"components"`
}

// RegisterRoutes registers routes for management endpoints.
func (managementController ManagementController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/management/health/readiness", managementController.readinessHandler).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", managementController.livenessHandler).Methods(http.MethodGet)
}

// readinessHandler handles the readiness endpoint.
// @Summary Get the readiness status
// @Description Get the readiness status of the service
// @Produce json
// @Success 200 {object} HealthResponse "Success response"
// @Router /management/health/readiness [get]
func (managementController ManagementController) readinessHandler(w http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{
		Status: "UP",
		Components: map[string]HealthStatus{
			"readinessState": {Status: "UP"},
		},
	}
	json.NewEncoder(w).Encode(response)
}

// livenessHandler handles the liveness endpoint.
// @Summary Get the liveness status
// @Description Get the liveness status of the service
// @Produce json
// @Success 200 {object} HealthResponse "Success response"
// @Router /management/health/liveness [get]
func (managementController ManagementController) livenessHandler(w http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{
		Status: "UP",
		Components: map[string]HealthStatus{
			"livenessState": {Status: "UP"},
		},
	}
	json.NewEncoder(w).Encode(response)
}
