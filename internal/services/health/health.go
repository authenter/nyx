package health

// Health are response for the api health endpoint
type Health struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
