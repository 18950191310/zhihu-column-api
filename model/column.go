package model

import "gorm.io/gorm"

// Column 专栏
type Column struct {
	gorm.Model
	Cover       string `gorm:"type:VARCHAR(300)" json:"cover"`
	Description string `gorm:"type:VARCHAR(300)" json:"description"`
	Title       string `gorm:"type:VARCHAR(100)" json:"title"`
	AuthorID    uint   `json:"author_id"`
}
