package health

import (
	"github.com/authenter/nyx/pkg/log"
	"net/http"
)

// Service is the object for Health Service
type Service struct {
	logger log.Logger
}

// NewService return new Service object
func NewService(logger log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// GetHealth return filled Health object
func (s *Service) GetHealth(serverState string) Health {
	var health Health
	switch serverState {
	case "created":
		health.Status = http.StatusAccepted
	case "starting":
		health.Status = http.StatusAccepted
	case "running":
		health.Status = http.StatusOK
	case "stopping":
		health.Status = http.StatusNoContent
	default:
		health.Status = http.StatusInternalServerError
	}

	health.Message = serverState
	return health
}
