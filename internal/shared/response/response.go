package response

type Meta struct {
	Page      int   `json:"page,omitempty"`
	Limit     int   `json:"limit,omitempty"`
	TotalData int64 `json:"total_data,omitempty"`
	TotalPage int   `json:"total_page,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    string        `json:"code"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    any            `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Meta    *Meta          `json:"meta,omitempty"`
}
