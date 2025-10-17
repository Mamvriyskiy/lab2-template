package models

type ValidationErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ErrorDescription struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
