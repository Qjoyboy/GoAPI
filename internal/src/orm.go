package src

type Task struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

type TaskRequest struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}
