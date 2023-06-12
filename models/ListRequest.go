package models

import "time"

type ListRequest struct {
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
	Keyword  string    `json:"keyword"`
	Lte      time.Time `json:"lte"`
	Gte      time.Time `json:"gte"`
}

func DefaultListResponse() ListRequest {
	return ListRequest{
		Page:     1,
		PageSize: 10,
		Keyword:  "",
		Lte:      time.Now().UTC(),
		Gte:      time.Now().UTC(),
	}
}
