package api_models

type ValidationError struct {
	Code   string         `json:"code"`
	Params map[string]any `json:"params,omitempty"`
}

type ValidationErrorsResponse struct {
	Errors map[string][]ValidationError `json:"errors"`
}
