package models

import "time"

type TodoItem struct {
	ID          uint       `gorm:"primaryKey;column:id" json:"id"`
	Title       string     `gorm:"size:255;not null;column:title" json:"title"`
	Description string     `gorm:"size:255;null;column:description" json:"description"`
	IsCompleted bool       `gorm:"default:false;column:isCompleted" json:"isCompleted"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	Notes       []TodoNote `gorm:"foreignKey:TodoItemID;constraint:OnDelete:CASCADE" json:"notes,omitempty"`
	UserID      uint       `gorm:"column:user_id; not null" json:"userId"`
	User        User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName overrides the table name used by TodoItem to `todos`
func (TodoItem) TableName() string {
	return "TodoItems"
}
