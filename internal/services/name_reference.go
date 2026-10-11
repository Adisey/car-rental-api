package services

import (
	"github.com/Adisey/car-rental-api/internal/api_models"
)

func IsEmptyNameReference(
	ref *api_models.NameReference,
) bool {
	return ref != nil &&
		ref.Id == nil &&
		ref.Name == nil
}
