package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetCarsService(ctx context.Context) ([]api_models.Car, error) {
	repo := repositories.NewCarRepository()
	dbCars, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	cars := make([]api_models.Car, 0, len(dbCars))
	for _, car := range dbCars {
		var bodyType *api_models.CodeReference
		if car.BodyType != nil {
			bodyTypeID := car.BodyType.ID.String()
			bodyTypeCode := car.BodyType.Code

			bodyType = &api_models.CodeReference{
				Id:   &bodyTypeID,
				Code: &bodyTypeCode,
			}
		}
		var brand *api_models.NameReference
		if car.Brand != nil {
			brandID := car.Brand.ID.String()
			brandCode := car.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandCode,
			}
		}
		var color *api_models.CarColor
		if car.Color != nil {
			colorID := car.Color.ID.String()
			colorCode := car.Color.Code
			colorHex := car.Color.Hex
			color = &api_models.CarColor{
				Id:   &colorID,
				Code: &colorCode,
				Hex:  colorHex,
			}
		}
		var driveType *api_models.CodeReference
		if car.DriveType != nil {
			driveID := car.DriveType.ID.String()
			driveCode := car.DriveType.Code
			driveType = &api_models.CodeReference{
				Id:   &driveID,
				Code: &driveCode,
			}
		}
		var fuelType *api_models.CodeReference
		if car.FuelType != nil {
			fuelID := car.FuelType.ID.String()
			fuelCode := car.FuelType.Code
			fuelType = &api_models.CodeReference{
				Id:   &fuelID,
				Code: &fuelCode,
			}
		}
		var model *api_models.Model
		if car.Model != nil {
			modelID := car.Model.ID.String()
			modelCode := car.Model.Name
			brandID := car.Model.Brand.ID.String()
			brandCode := car.Model.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandCode,
			}
			model = &api_models.Model{
				Id:    modelID,
				Name:  modelCode,
				Brand: brand,
			}
		}
		var transmissionType *api_models.CodeReference
		if car.TransmissionType != nil {
			ID := car.TransmissionType.ID.String()
			Code := car.TransmissionType.Code
			transmissionType = &api_models.CodeReference{
				Id:   &ID,
				Code: &Code,
			}
		}
		cars = append(cars, api_models.Car{
			Id:               car.ID.String(),
			Name:             car.Name,
			Description:      car.Description,
			PowerHp:          car.PowerHp,
			BodyType:         bodyType,
			Brand:            brand,
			Color:            color,
			DriveType:        driveType,
			FuelType:         fuelType,
			Model:            model,
			TransmissionType: transmissionType,
			CreatedAt:        car.CreatedAt,
			UpdatedAt:        car.UpdatedAt,
		})
	}
	return cars, nil
}

func GetCarByIDService(
	ctx context.Context,
	id string,
) (*api_models.Car, error) {
	repo := repositories.NewCarRepository()
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbCar, err := repo.GetByIDRepository(ctx, carID)
	if err != nil {
		return nil, err
	}
	var bodyType *api_models.CodeReference
	if dbCar.BodyType != nil {
		bodyTypeID := dbCar.BodyType.ID.String()
		bodyTypeCode := dbCar.BodyType.Code
		bodyType = &api_models.CodeReference{
			Id:   &bodyTypeID,
			Code: &bodyTypeCode,
		}
	}
	var brand *api_models.NameReference
	if dbCar.Brand != nil {
		brandID := dbCar.Brand.ID.String()
		brandCode := dbCar.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandCode,
		}
	}
	var color *api_models.CarColor
	if dbCar.Color != nil {
		colorID := dbCar.Color.ID.String()
		colorCode := dbCar.Color.Code
		colorHex := dbCar.Color.Hex
		color = &api_models.CarColor{
			Id:   &colorID,
			Code: &colorCode,
			Hex:  colorHex,
		}
	}
	var driveType *api_models.CodeReference
	if dbCar.DriveType != nil {
		itemID := dbCar.DriveType.ID.String()
		itemCode := dbCar.DriveType.Code
		driveType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	var fuelType *api_models.CodeReference
	if dbCar.FuelType != nil {
		itemID := dbCar.FuelType.ID.String()
		itemCode := dbCar.FuelType.Code
		fuelType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	var model *api_models.Model
	if dbCar.Model != nil {
		modelID := dbCar.Model.ID.String()
		modelCode := dbCar.Model.Name
		brandID := dbCar.Model.Brand.ID.String()
		brandCode := dbCar.Model.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandCode,
		}
		model = &api_models.Model{
			Id:    modelID,
			Name:  modelCode,
			Brand: brand,
		}
	}
	var transmissionType *api_models.CodeReference
	if dbCar.TransmissionType != nil {
		itemID := dbCar.TransmissionType.ID.String()
		itemCode := dbCar.TransmissionType.Code
		transmissionType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	return &api_models.Car{
		Id:               dbCar.ID.String(),
		Name:             dbCar.Name,
		Description:      dbCar.Description,
		PowerHp:          dbCar.PowerHp,
		BodyType:         bodyType,
		Brand:            brand,
		Color:            color,
		DriveType:        driveType,
		FuelType:         fuelType,
		Model:            model,
		TransmissionType: transmissionType,
		CreatedAt:        dbCar.CreatedAt,
		UpdatedAt:        dbCar.UpdatedAt,
	}, nil
}

func CreateCarService(
	ctx context.Context,
	request api_models.CreateCarRequest,
) (*api_models.Car, error) {
	validationResult := validation.ValidateCreateCarRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}

	var bodyTypeID *uuid.UUID
	if request.BodyType != nil &&
		!IsEmptyCodeReference(request.BodyType) {
		var err error
		bodyTypeID, err = ResolveBodyTypeID(
			ctx,
			request.BodyType,
		)
		if err != nil {
			return nil, err
		}
	}
	var brandID *uuid.UUID
	if request.Brand != nil &&
		!IsEmptyBrand(request.Brand) {
		var err error
		brandID, err = ResolveBrandID(
			ctx,
			request.Brand,
		)
		if err != nil {
			return nil, err
		}
	}
	var colorID *uuid.UUID
	if request.Color != nil &&
		!IsEmptyColor(request.Color) {
		var err error
		colorID, err = ResolveColorID(
			ctx,
			request.Color,
		)
		if err != nil {
			return nil, err
		}
	}
	var driveTypeID *uuid.UUID
	if request.DriveType != nil &&
		!IsEmptyCodeReference(request.DriveType) {
		var err error
		driveTypeID, err = ResolveDriveTypeID(
			ctx,
			request.DriveType,
		)
		if err != nil {
			return nil, err
		}
	}
	var FuelTypeID *uuid.UUID
	if request.FuelType != nil &&
		!IsEmptyCodeReference(request.FuelType) {
		var err error
		FuelTypeID, err = ResolveFuelTypeID(
			ctx,
			request.FuelType,
		)
		if err != nil {
			return nil, err
		}
	}
	var modelID *uuid.UUID
	if request.Model != nil {
		model := request.Model
		if model.Brand == nil &&
			brandID != nil {
			brandIDString := brandID.String()
			model.Brand = &api_models.NameReference{
				Id: &brandIDString,
			}
		}
		if !IsEmptyModel(model) {
			var err error
			modelID, err = ResolveModelID(
				ctx,
				model,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var transmissionTypeID *uuid.UUID
	if request.TransmissionType != nil &&
		!IsEmptyCodeReference(request.TransmissionType) {
		var err error
		transmissionTypeID, err = ResolveTransmissionTypeID(
			ctx,
			request.TransmissionType,
		)
		if err != nil {
			return nil, err
		}
	}

	repo := repositories.NewCarRepository()
	car := &db_models.Car{
		Name:           request.Name,
		Description:    request.Description,
		BodyTypeID:     bodyTypeID,
		PowerHp:        request.PowerHp,
		BrandID:        brandID,
		ColorID:        colorID,
		DriveID:        driveTypeID,
		FuelID:         FuelTypeID,
		ModelID:        modelID,
		TransmissionID: transmissionTypeID,
	}
	err := repo.CreateRepository(ctx, car)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)

	if err != nil {
		return nil, err
	}
	var color *api_models.CarColor
	if car.Color != nil {
		colorID := car.Color.ID.String()
		colorCode := car.Color.Code
		colorHex := car.Color.Hex
		color = &api_models.CarColor{
			Id:   &colorID,
			Code: &colorCode,
			Hex:  colorHex,
		}
	}
	var bodyType *api_models.CodeReference
	if car.BodyType != nil {
		bodyTypeID := car.BodyType.ID.String()
		bodyTypeCode := car.BodyType.Code
		bodyType = &api_models.CodeReference{
			Id:   &bodyTypeID,
			Code: &bodyTypeCode,
		}
	}
	var brand *api_models.NameReference
	if car.Brand != nil {
		brandID := car.Brand.ID.String()
		brandName := car.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	var driveType *api_models.CodeReference
	if car.DriveType != nil {
		itemID := car.DriveType.ID.String()
		itemCode := car.DriveType.Code
		driveType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	var fuelType *api_models.CodeReference
	if car.FuelType != nil {
		itemID := car.FuelType.ID.String()
		itemCode := car.FuelType.Code
		fuelType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	var model *api_models.Model
	if car.Model != nil {
		var modelBrand *api_models.NameReference
		if car.Model.Brand != nil {
			brandID := car.Model.Brand.ID.String()
			brandName := car.Model.Brand.Name
			modelBrand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}
		newModelID := car.Model.ID.String()
		newModelName := car.Model.Name
		model = &api_models.Model{
			Id:    newModelID,
			Name:  newModelName,
			Brand: modelBrand,
		}
	}
	var transmissionType *api_models.CodeReference
	if car.TransmissionType != nil {
		itemID := car.TransmissionType.ID.String()
		itemCode := car.TransmissionType.Code
		transmissionType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	return &api_models.Car{
		Id:               car.ID.String(),
		Name:             car.Name,
		Description:      car.Description,
		PowerHp:          car.PowerHp,
		Color:            color,
		BodyType:         bodyType,
		Brand:            brand,
		DriveType:        driveType,
		FuelType:         fuelType,
		Model:            model,
		TransmissionType: transmissionType,
		CreatedAt:        car.CreatedAt,
		UpdatedAt:        car.UpdatedAt,
	}, nil
}

func UpdateCarService(
	ctx context.Context,
	id string,
	request api_models.UpdateCarRequest,
) (*api_models.Car, error) {
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	validationResult := validation.ValidateUpdateCarRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	var (
		bodyID     *uuid.UUID
		updateBody bool
	)
	if request.BodyType != nil {
		updateBody = true
		if !IsEmptyCodeReference(request.BodyType) {
			bodyID, err = ResolveBodyTypeID(
				ctx,
				request.BodyType,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		brandID     *uuid.UUID
		updateBrand bool
	)
	if request.Brand != nil {
		updateBrand = true
		if !IsEmptyBrand(request.Brand) {
			brandID, err = ResolveBrandID(
				ctx,
				request.Brand,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		colorID     *uuid.UUID
		updateColor bool
	)
	if request.Color != nil {
		updateColor = true
		if !IsEmptyColor(request.Color) {
			colorID, err = ResolveColorID(
				ctx,
				request.Color,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		driveID     *uuid.UUID
		updateDrive bool
	)
	if request.DriveType != nil {
		updateDrive = true
		if !IsEmptyCodeReference(request.DriveType) {
			driveID, err = ResolveDriveTypeID(
				ctx,
				request.DriveType,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		fuelID     *uuid.UUID
		updateFuel bool
	)
	if request.FuelType != nil {
		updateFuel = true
		if !IsEmptyCodeReference(request.FuelType) {
			fuelID, err = ResolveFuelTypeID(
				ctx,
				request.FuelType,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		modelID     *uuid.UUID
		updateModel bool
	)
	if request.Model != nil {
		updateModel = true
		model := request.Model
		if model.Brand == nil &&
			brandID != nil {
			brandIDString := brandID.String()
			model.Brand = &api_models.NameReference{
				Id: &brandIDString,
			}
		}
		if !IsEmptyModel(model) {
			var err error
			modelID, err = ResolveModelID(
				ctx,
				model,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		transmissionID     *uuid.UUID
		updateTransmission bool
	)
	if request.TransmissionType != nil {
		updateTransmission = true
		if !IsEmptyCodeReference(request.TransmissionType) {
			transmissionID, err = ResolveTransmissionTypeID(
				ctx,
				request.TransmissionType,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	relations := repositories.CarRelationsToUpdate{
		BodyID:             bodyID,
		UpdateBody:         &updateBody,
		BrandID:            brandID,
		UpdateBrand:        &updateBrand,
		ColorID:            colorID,
		UpdateColor:        &updateColor,
		DriveID:            driveID,
		UpdateDrive:        &updateDrive,
		FuelID:             fuelID,
		UpdateFuel:         &updateFuel,
		ModelID:            modelID,
		UpdateModel:        &updateModel,
		TransmissionID:     transmissionID,
		UpdateTransmission: &updateTransmission,
	}
	repo := repositories.NewCarRepository()
	car, err := repo.UpdateRepository(
		ctx,
		carID,
		request,
		relations,
	)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)
	if err != nil {
		return nil, err
	}
	var bodyType *api_models.CodeReference
	if car.BodyType != nil {
		bodyTypeID := car.BodyType.ID.String()
		bodyTypeCode := car.BodyType.Code
		bodyType = &api_models.CodeReference{
			Id:   &bodyTypeID,
			Code: &bodyTypeCode,
		}
	}
	var brand *api_models.NameReference
	if car.Brand != nil {
		brandID := car.Brand.ID.String()
		brandName := car.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	var color *api_models.CarColor
	if car.Color != nil {
		colorID := car.Color.ID.String()
		colorCode := car.Color.Code
		color = &api_models.CarColor{
			Id:   &colorID,
			Code: &colorCode,
		}
	}
	var driveType *api_models.CodeReference
	if car.DriveType != nil {
		itemID := car.DriveType.ID.String()
		itemCode := car.DriveType.Code
		driveType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}
	var model *api_models.Model
	if car.Model != nil {
		var modelBrand *api_models.NameReference
		if car.Model.Brand != nil {
			brandID := car.Model.Brand.ID.String()
			brandName := car.Model.Brand.Name
			modelBrand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}
		newModelID := car.Model.ID.String()
		newModelName := car.Model.Name
		model = &api_models.Model{
			Id:    newModelID,
			Name:  newModelName,
			Brand: modelBrand,
		}
	}
	var transmissionType *api_models.CodeReference
	if car.TransmissionType != nil {
		itemID := car.TransmissionType.ID.String()
		itemCode := car.TransmissionType.Code
		transmissionType = &api_models.CodeReference{
			Id:   &itemID,
			Code: &itemCode,
		}
	}

	return &api_models.Car{
		Id:               car.ID.String(),
		Name:             car.Name,
		Description:      car.Description,
		PowerHp:          car.PowerHp,
		BodyType:         bodyType,
		Brand:            brand,
		Color:            color,
		DriveType:        driveType,
		Model:            model,
		TransmissionType: transmissionType,
		CreatedAt:        car.CreatedAt,
		UpdatedAt:        car.UpdatedAt,
	}, nil
}

func DeleteCarService(ctx context.Context, id string) error {
	carID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	repo := repositories.NewCarRepository()
	return repo.DeleteRepository(
		ctx,
		carID,
	)
}
