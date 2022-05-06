package v1

import (
	"net/http"

	"template-service/api/rest"
	"template-service/services"
)

func NewHealthController(services services.Services) *HealthController {
	return &HealthController{
		services: services,
	}
}

type HealthController struct {
	services services.Services
}

// GetHealth handles a health request
func (c *HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {

	// Get the context
	ctx := r.Context()

	statusCode := http.StatusOK
	healthStatus := c.services.Health.GetHealth(ctx)

	if !healthStatus.Healthy {
		statusCode = http.StatusServiceUnavailable
	}

	rest.RespondWithJSON(ctx, w, statusCode, healthStatus)
}
