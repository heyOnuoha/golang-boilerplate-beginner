package models

import "time"

type TodoNote struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	TodoItemID uint      `gorm:"column:todoItemId;not null" json:"todoItemId"`
	Note       string    `gorm:"column:note;not null" json:"note"`
	CreatedAt  time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	TodoItem   *TodoItem `gorm:"foreignKey:TodoItemID;references:ID" json:"-"` // The reference to parent TodoItem, but excluded from JSON
}

func (TodoNote) TableName() string {
	return "TodoNotes"
}
