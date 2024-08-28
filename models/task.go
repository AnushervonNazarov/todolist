package models

type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      string `json:"isdone"`
	// IsDeleted   bool `gorm:"default:false"`
}

type DefaultResponse struct {
	Message string `json:"message"`
}
