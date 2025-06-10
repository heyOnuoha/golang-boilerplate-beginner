package models

type User struct {
	ID           uint       `gorm:"primaryKey;column:id" json:"id"`
	Email        string     `gorm:"column:email;not null;unique" json:"email"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	PasswordHash string     `gorm:"column:passwordHash;not null" json:"-"` // Using json:"-" to exclude from JSON responses
	TodoItems    []TodoItem `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"todoItems,omitempty"`
}

func (User) TableName() string {
	return "Users"
}
