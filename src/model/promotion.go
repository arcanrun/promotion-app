package model

import (
	"github.com/google/uuid"
	"time"
)

type Promotion struct {
	Id             int64
	OriginalId     uuid.UUID `gorm:"type:uuid;not null;unique"`
	Price          float64   `gorm:"type:numeric;not null"`
	ExpirationDate time.Time `gorm:"type:timestamp with time zone;not null"`
}
