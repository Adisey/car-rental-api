package car_brand

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func GetBrandsService(ctx context.Context) ([]api_models.Brand, error) {
	repo := NewBrandRepository()

	dbBrands, err := repo.GetBrandAll(ctx)
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

	brandID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	repo := NewBrandRepository()

	dbBrand, err := repo.GetBrandByIDRepository(ctx, brandID)
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
	request api_models.CreateNameEntityRequest,
) (*api_models.Brand, error) {

	validationResult := shared_validation.ValidateCreateNameRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	repo := NewBrandRepository()

	existingBrand, errExist := repo.GetBrandByNameRepository(
		ctx,
		request.Name,
	)

	if errExist == nil {
		return &api_models.Brand{
			Id:   existingBrand.ID.String(),
			Name: existingBrand.Name,
		}, nil
	}

	brand := &BrandDBModel{
		Name: request.Name,
	}

	errNew := repo.CreateBrandRepository(ctx, brand)
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
	request api_models.UpdateNameEntityRequest,
) (*api_models.Brand, error) {

	repo := NewBrandRepository()

	brandID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}

	validationResult := shared_validation.ValidateUpdateNameRequest(
		request,
	)

	if request.Name != nil {

		existingBrand, err := repo.GetBrandByNameRepository(
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

	brand, errUpdate := repo.UpdateBrandRepository(
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
