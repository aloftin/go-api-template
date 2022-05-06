package repositories

import (
	"context"
	"database/sql"

	"template-service/models"

	// Import the DB driver here so that DB concerns are encapsulated
	_ "github.com/lib/pq"
)

// HealthRepository interface definition
type HealthRepository interface {
	GetHealth(ctx context.Context) models.ComponentHealthStatus
}

type healthRepository struct {
	db *sql.DB
}

// NewHealthRepository returns an instance of HealthRepository
func NewHealthRepository(db *sql.DB) HealthRepository {
	return &healthRepository{
		db: db,
	}
}

func (r *healthRepository) GetHealth(ctx context.Context) models.ComponentHealthStatus {
	status := models.ComponentHealthStatus{
		Component: "Database",
		Healthy:   true,
	}

	if err := r.db.Ping(); err != nil {
		status.Healthy = false
		status.Error = err.Error()
	}

	return status
}
