package service

import (
	"github.com/promotionApp/src/dto"
	"github.com/promotionApp/src/model"
	"github.com/promotionApp/src/repository"
	"math/rand"
)

func GetAllPromotions() []model.Promotion {
	return repository.GetAllPromotions()
}

func AddPromotion(dto dto.PromotionDto) {
	promotion := model.Promotion{
		Id:             rand.Int63(),
		OriginalId:     dto.OriginalId,
		Price:          dto.Price,
		ExpirationDate: dto.ExpirationDate,
	}

	repository.Save(promotion)
}

func GetPromotingById(id int64) (dto.PromotionResponseDto, error) {
	if promotion, err := repository.GetPromotingById(id); err != nil {
		return dto.PromotionResponseDto{}, err
	} else {
		return dto.PromotionResponseDto{
			Id:             promotion.Id,
			OriginalId:     promotion.OriginalId,
			Price:          promotion.Price,
			ExpirationDate: promotion.ExpirationDate,
		}, nil
	}
}
