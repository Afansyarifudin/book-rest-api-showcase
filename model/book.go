package model

import "time"

// import "time"

// Book represent the model for an book
type Book struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Namebook  string    `json:"name_book" gorm:"type:varchar(50)"`
	Author    string    `json:"author" gorm:"type:varchar(20)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
