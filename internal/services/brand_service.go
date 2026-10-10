package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetBrandsService(ctx context.Context) ([]api_models.Brand, error) {
	repo := repositories.NewBrandRepository()

	dbBrands, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	brands := make([]api_models.Brand, 0, len(dbBrands))

	for _, brand := range dbBrands {
		brands = append(brands, api_models.Brand{
			Id:   brand.ID.String(),
			Name: brand.Name,
		})
	}

	return brands, nil
}

func GetBrandByIDService(
	ctx context.Context,
	id string,
) (*api_models.Brand, error) {
	repo := repositories.NewBrandRepository()

	brandID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbBrand, err := repo.GetByIDRepository(ctx, brandID)
	if err != nil {
		return nil, err
	}

	return &api_models.Brand{
		Id:   dbBrand.ID.String(),
		Name: dbBrand.Name,
	}, nil
}

func CreateBrandService(
	ctx context.Context,
	request api_models.CreateBrandRequest,
) (*api_models.Brand, error) {

	validationResult := validation.ValidateCreateBrandRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	repo := repositories.NewBrandRepository()

	existingBrand, errExist := repo.GetByNameRepository(
		ctx,
		request.Name,
	)

	if errExist == nil {
		return &api_models.Brand{
			Id:   existingBrand.ID.String(),
			Name: existingBrand.Name,
		}, nil
	}

	brand := &db_models.Brand{
		Name: request.Name,
	}

	errNew := repo.CreateRepository(ctx, brand)
	if errNew != nil {
		return nil, errNew
	}

	return &api_models.Brand{
		Id:   brand.ID.String(),
		Name: brand.Name,
	}, nil
}

func UpdateBrandService(
	ctx context.Context,
	id string,
	request api_models.UpdateBrandRequest,
) (*api_models.Brand, error) {

	repo := repositories.NewBrandRepository()

	brandID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}

	validationResult := validation.ValidateUpdateBrandRequest(
		request,
	)

	if request.Name != nil {

		existingBrand, err := repo.GetByNameRepository(
			ctx,
			*request.Name,
		)

		if err == nil &&
			existingBrand.ID != brandID {

			validationResult.Add(
				"code",
				"already_exists",
				nil,
			)
		}
	}

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	brand, errUpdate := repo.UpdateRepository(
		ctx,
		brandID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return &api_models.Brand{
		Id:   brand.ID.String(),
		Name: brand.Name,
	}, nil
}
