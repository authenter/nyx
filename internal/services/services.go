package services

import "github.com/authenter/nyx/internal/services/health"

// Services holds all available Services
type Services struct {
	Health *health.Service
}
