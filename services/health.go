package services

import (
	"context"
	"database/sql"

	"template-service/models"
	"template-service/repositories"
)

// HealthService interface definition
type HealthService interface {
	GetHealth(ctx context.Context) models.HealthStatus
}

type healthService struct {
	healthRepo repositories.HealthRepository
}

// NewHealthService returns an instance of HealthService
func NewHealthService(db *sql.DB) HealthService {
	return &healthService{healthRepo: repositories.NewHealthRepository(db)}
}

func (s *healthService) GetHealth(ctx context.Context) models.HealthStatus {
	// Check the health of the repository
	// repositoryHealth := s.healthRepo.GetHealth(ctx) // TODO
	repositoryHealth := models.ComponentHealthStatus{
		Component: "stubbed-database",
		Healthy:   true,
	}

	healthStatus := models.HealthStatus{
		Healthy:          true,
		ComponentsHealth: []models.ComponentHealthStatus{repositoryHealth},
	}

	if !repositoryHealth.Healthy {
		// Set the overall health to false
		healthStatus.Healthy = false
		healthStatus.Error = "Repository unavailable"
	}

	return healthStatus
}
