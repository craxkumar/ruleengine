package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ManagementController handles management related endpoints.
type ManagementController struct {
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
// @Success 200
// @Router /management/health/readiness [get]
func (managementController ManagementController) readinessHandler(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"readinessState": map[string]interface{}{"status": "UP"}}})
}

// livenessHandler handles the liveness endpoint.
// @Summary Get the liveness status
// @Description Get the liveness status of the service
// @Produce json
// @Success 200
// @Router /management/health/liveness [get]
func (managementController ManagementController) livenessHandler(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP", "components": map[string]interface{}{"livenessState": map[string]interface{}{"status": "UP"}}})
}
