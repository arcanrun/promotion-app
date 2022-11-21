package model

import (
	"github.com/google/uuid"
	"time"
)

type Promotion struct {
	Id             int64
	OriginalId     uuid.UUID
	Price          float64
	ExpirationDate time.Time
}
