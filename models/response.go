package models

type Responses struct {
	Responses []HttpResponse `json:"responses"`
}

type HttpResponse struct {
	StatusCode int            `json:"status_code"`
	Method     string         `json:"method"`
	Body       map[string]any `json:"body"`
	Route      string         `json:"route"`
}
