package handler

type SuccessResponse struct {
	Message map[string]string `json:"message,omitempty" example:"notification sent successfully!"`
}
