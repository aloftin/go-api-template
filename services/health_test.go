package services

import (
	"context"
	"template-service/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	healthService := NewHealthService(nil)

	repositoryHealth := models.ComponentHealthStatus{
		Component: "stubbed-database",
		Healthy:   true,
	}

	want := models.HealthStatus{
		Healthy:          true,
		ComponentsHealth: []models.ComponentHealthStatus{repositoryHealth},
	}

	got := healthService.GetHealth(context.Background())
	assert.Equal(t, want, got)
}
