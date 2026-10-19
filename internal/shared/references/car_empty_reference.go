package shared_references

import (
	"github.com/Adisey/car-rental-api/internal/api_models"
)

func IsEmptyCodeReference(bodyType *api_models.CodeReference) bool {
	return bodyType != nil &&
		bodyType.Id == nil &&
		bodyType.Code == nil
}

func IsEmptyNameReference(brand *api_models.NameReference) bool {
	return brand != nil &&
		brand.Id == nil &&
		brand.Name == nil
}

func IsEmptyColorReference(color *api_models.CarColor) bool {
	return color != nil &&
		color.Id == nil &&
		color.Code == nil
	// hex can be empty
}

func IsEmptyModelReference(
	model *api_models.CarModel,
) bool {
	return model != nil &&
		model.Id == nil &&
		model.Name == nil &&
		model.Brand == nil
}
