package models

type Tasks struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	TotalTime   string `json:"total_time"`
	Done        bool   `json:"done"`
}
