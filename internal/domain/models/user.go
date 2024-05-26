package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	UUID      uuid.UUID `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
