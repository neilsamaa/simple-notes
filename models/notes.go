package models

import "time"

type Notes struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `json:"content"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
