package models

// HealthStatus holds results of a health check
type HealthStatus struct {
	Healthy          bool                    `json:"healthy"`
	Error            string                  `json:"error,omitempty"`
	ComponentsHealth []ComponentHealthStatus `json:"componentsHealth"`
}

// ComponentHealthStatus holds results of a component health check
type ComponentHealthStatus struct {
	Component string `json:"component"`
	Healthy   bool   `json:"healthy"`
	Error     string `json:"error,omitempty"`
}
