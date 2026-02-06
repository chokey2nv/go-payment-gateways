package client

type responseEnvelope struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    *Meta  `json:"meta,omitempty"`
}

type Meta struct {
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	PerPage   int    `json:"perPage"`
	Total     int    `json:"total"`
	Skipped   int    `json:"skipped"`
	Page      int    `json:"page"`
	PageCount int    `json:"pageCount"`
}
