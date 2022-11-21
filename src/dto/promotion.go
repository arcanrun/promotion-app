package dto

import (
	"github.com/google/uuid"
	"time"
)

type PromotionDto struct {
	OriginalId     uuid.UUID
	Price          float64
	ExpirationDate time.Time
}

type PromotionResponseDto struct {
	Id             int64
	OriginalId     uuid.UUID
	Price          float64
	ExpirationDate time.Time
}
