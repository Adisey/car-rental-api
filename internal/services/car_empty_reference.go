package services

import (
	"github.com/Adisey/car-rental-api/internal/api_models"
)

func IsEmptyCodeReference(bodyType *api_models.CodeReference) bool {
	return bodyType != nil &&
		bodyType.Id == nil &&
		bodyType.Code == nil
}

func IsEmptyBrand(brand *api_models.NameReference) bool {
	return brand != nil &&
		brand.Id == nil &&
		brand.Name == nil
}
