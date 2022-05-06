package services

import (
	"database/sql"
	"os"
	"strconv"
)

// ServicesConfig function definition
type ServicesConfig func(*Services) error

// Services defines all the services for this application
type Services struct {
	Health HealthService
	db     *sql.DB
}

// NewServices configures all services for the app and returns them as properties on Services
func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

// WithDB creates a DB connection for the app
func WithDB(dialect string, connString string) ServicesConfig {
	return func(s *Services) error {
		db, err := sql.Open(dialect, connString)
		dbConns := 75

		if os.Getenv("MAX_DB_CONNS") != "" {
			env, err := strconv.Atoi(os.Getenv("MAX_DB_CONNS"))

			if err == nil {
				dbConns = env
			}
		}

		db.SetMaxOpenConns(dbConns)
		db.SetMaxIdleConns(dbConns)

		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

// WithHealth creates an instance of HealthService and assigns it to Services
func WithHealth() ServicesConfig {
	return func(s *Services) error {
		s.Health = NewHealthService(s.db)
		return nil
	}
}

// Close loses the database connection
func (s *Services) Close() error {
	return s.db.Close()
}
