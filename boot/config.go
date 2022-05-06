package boot

import (
	"fmt"
	"os"
)

const (
	// DefaultPort is the default HTTP port
	DefaultPort = "80"

	// DBType is the type of database
	DBType = "postgres"

	// DefaultEnableTLS is the default setting whether to enable TLS encryption
	DefaultEnableTLS = "false"
)

// GetConnectionString returns the formatted DB connection string
func GetConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		DBType,
		os.Getenv("TS_PG_USER"),
		os.Getenv("TS_PG_PW"),
		os.Getenv("TS_PG_HOST"),
		os.Getenv("TS_PG_DB"))
}
