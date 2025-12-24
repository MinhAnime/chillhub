package response

type Envelope struct {
	Success bool       `json:"success"`
	Status  int        `json:"status,omitempty"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorBody `json:"error,omitempty"`
}

type ErrorBody struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}
