package src

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

type TaskRequest struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}
