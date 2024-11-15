package models

type Store struct {
	StoreID   string `json:"store_id"`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}

type Job struct {
	Status string  `json:"status"`
	JobID  string  `json:"job_id"`
	Errors []Error `json:"error,omitempty"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type Error struct {
	StoreID string `json:"store_id"`
	Message string `json:"error"`
}

type SubmitRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}
