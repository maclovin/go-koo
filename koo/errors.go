package koo

type APIError struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
