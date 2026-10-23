package car

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	car_body "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
	car_drive "github.com/Adisey/car-rental-api/internal/cars/dictionaries/drive_type"
	car_fuel "github.com/Adisey/car-rental-api/internal/cars/dictionaries/fuel_type"
	car_model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
	car_transmission "github.com/Adisey/car-rental-api/internal/cars/dictionaries/transmission_type"
	shared_references "github.com/Adisey/car-rental-api/internal/shared/references"
)

func GetCarsService(ctx context.Context) ([]api_models.Car, error) {
	repo := NewCarRepository()
	dbCars, err := repo.GetCarAllRepository(ctx)
	if err != nil {
		return nil, err
	}
	cars := make([]api_models.Car, 0, len(dbCars))
	// for _, car := range dbCars {
	// 	cars = append(cars, mapCarToAPI(car))
	// }

	for i := range dbCars {
		cars = append(cars, mapCarToAPI(&dbCars[i]))
	}
	return cars, nil
}

func GetCarByIDService(
	ctx context.Context,
	id string,
) (*api_models.Car, error) {
	repo := NewCarRepository()
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbCar, err := repo.GetCarByIDRepository(ctx, carID)
	if err != nil {
		return nil, err
	}
	apiCar := mapCarToAPI(dbCar)
	return &apiCar, nil
}

func CreateCarService(
	ctx context.Context,
	request api_models.CreateCarRequest,
) (*api_models.Car, error) {
	validationResult := ValidateCreateCarRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	println("service-3->", "isValid")

	var bodyTypeID *uuid.UUID
	if request.BodyType != nil &&
		!shared_references.IsEmptyCodeReference(request.BodyType) {
		var err error
		bodyTypeID, err = car_body.ResolveBodyTypeID(
			ctx,
			request.BodyType,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-bodyTypeID->", bodyTypeID)
	var brandID *uuid.UUID
	if request.Brand != nil &&
		!shared_references.IsEmptyNameReference(request.Brand) {
		var err error
		brandID, err = car_brand.ResolveBrandID(
			ctx,
			request.Brand,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-brandID->", brandID)
	var colorID *uuid.UUID
	if request.Color != nil &&
		!shared_references.IsEmptyColorReference(request.Color) {
		var err error
		colorID, err = car_color.ResolveColorID(
			ctx,
			request.Color,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-colorID->", colorID)
	var driveTypeID *uuid.UUID
	if request.DriveType != nil &&
		!shared_references.IsEmptyCodeReference(request.DriveType) {
		var err error
		driveTypeID, err = car_drive.ResolveDriveTypeID(
			ctx,
			request.DriveType,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-driveTypeID->", driveTypeID)
	var FuelTypeID *uuid.UUID
	if request.FuelType != nil &&
		!shared_references.IsEmptyCodeReference(request.FuelType) {
		var err error
		FuelTypeID, err = car_fuel.ResolveFuelTypeID(
			ctx,
			request.FuelType,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-FuelTypeID->", FuelTypeID)
	var modelID *uuid.UUID
	if request.Model != nil {
		model := request.Model
		if model.Brand == nil &&
			brandID != nil {
			brandIDString := brandID.String()
			println("service-modelID-brandIDString->", brandIDString)
			model.Brand = &api_models.NameReference{
				Id: &brandIDString,
			}
		}
		if !shared_references.IsEmptyModelReference(model) {
			var err error
			println("service-modelID------>", 7)
			modelID, err = car_model.ResolveModelID(
				ctx,
				model,
			)
			println("service-modelID------>", 8)
			println("service-modelID--id-->", modelID)
			println("service-modelID--err->", err)
			if err != nil {
				return nil, err
			}
		}
	}
	println("service-modelID->", modelID)
	var transmissionTypeID *uuid.UUID
	if request.TransmissionType != nil &&
		!shared_references.IsEmptyCodeReference(request.TransmissionType) {
		var err error
		transmissionTypeID, err = car_transmission.ResolveTransmissionTypeID(
			ctx,
			request.TransmissionType,
		)
		if err != nil {
			return nil, err
		}
	}
	println("service-transmissionTypeID->", transmissionTypeID)

	repo := NewCarRepository()
	car := &CarDBModel{
		Name:           request.Name,
		Description:    request.Description,
		PowerHp:        request.PowerHp,
		BodyTypeID:     bodyTypeID,
		BrandID:        brandID,
		ColorID:        colorID,
		DriveID:        driveTypeID,
		FuelID:         FuelTypeID,
		ModelID:        modelID,
		TransmissionID: transmissionTypeID,
	}
	println("service-car->", car)
	err := repo.CreateCarRepository(ctx, car)
	if err != nil {
		return nil, err
	}
	dbCar, err := repo.GetCarByIDRepository(ctx, car.ID)
	if err != nil {
		return nil, err
	}
	apiCar := mapCarToAPI(dbCar)
	return &apiCar, nil
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
	validationResult := ValidateUpdateCarRequest(
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
		if !shared_references.IsEmptyCodeReference(request.BodyType) {
			bodyID, err = car_body.ResolveBodyTypeID(
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
		if !shared_references.IsEmptyNameReference(request.Brand) {
			brandID, err = car_brand.ResolveBrandID(
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
		if !shared_references.IsEmptyColorReference(request.Color) {
			colorID, err = car_color.ResolveColorID(
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
		if !shared_references.IsEmptyCodeReference(request.DriveType) {
			driveID, err = car_drive.ResolveDriveTypeID(
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
		if !shared_references.IsEmptyCodeReference(request.FuelType) {
			fuelID, err = car_fuel.ResolveFuelTypeID(
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
		if !shared_references.IsEmptyModelReference(model) {
			var err error
			modelID, err = car_model.ResolveModelID(
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
		if !shared_references.IsEmptyCodeReference(request.TransmissionType) {
			transmissionID, err = car_transmission.ResolveTransmissionTypeID(
				ctx,
				request.TransmissionType,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	relations := CarRelationsToUpdate{
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
	repo := NewCarRepository()
	car, err := repo.UpdateCarRepository(
		ctx,
		carID,
		request,
		relations,
	)
	if err != nil {
		return nil, err
	}
	dbCar, err := repo.GetCarByIDRepository(ctx, car.ID)
	if err != nil {
		return nil, err
	}
	apiCar := mapCarToAPI(dbCar)
	return &apiCar, nil

}

func DeleteCarService(ctx context.Context, id string) error {
	carID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	repo := NewCarRepository()
	return repo.DeleteCarRepository(
		ctx,
		carID,
	)
}

func mapCarToAPI(car *CarDBModel) api_models.Car {
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
		modelName := car.Model.Name

		var modelBrand *api_models.NameReference
		if car.Model.Brand != nil {
			brandID := car.Model.Brand.ID.String()
			brandName := car.Model.Brand.Name
			modelBrand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}

		model = &api_models.Model{
			Id:    modelID,
			Name:  modelName,
			Brand: modelBrand,
		}
	}

	var transmissionType *api_models.CodeReference
	if car.TransmissionType != nil {
		id := car.TransmissionType.ID.String()
		code := car.TransmissionType.Code
		transmissionType = &api_models.CodeReference{
			Id:   &id,
			Code: &code,
		}
	}

	return api_models.Car{
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
	}
}
