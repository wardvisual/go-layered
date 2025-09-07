package model

type TaskRequest struct {
	Title string `json:"title" validate:"required"`
	Done  *bool  `json:"done,omitempty"`
}

type TaskResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
