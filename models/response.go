package models

// Response standard for all request.
type Response struct {
	Success bool
	Message string
	Data    interface{}
}
