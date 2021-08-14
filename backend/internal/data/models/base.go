package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (b *BaseModel) BeforeCreate(_ *gorm.DB) error {
	b.ID = uuid.NewV4()
	return nil
}
