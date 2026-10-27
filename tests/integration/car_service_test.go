//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"

	car "github.com/Adisey/car-rental-api/internal/cars"
)

func Test_Car_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	bodyCode := fmt.Sprintf(
		"SEDAN_%d",
		time.Now().UnixNano(),
	)

	brandName := fmt.Sprintf(
		"TOYOTA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"CAMRY_%d",
		time.Now().UnixNano(),
	)

	colorCode := fmt.Sprintf(
		"BLACK_%d",
		time.Now().UnixNano(),
	)

	colorHex := "#000000"

	driveCode := fmt.Sprintf(
		"FWD_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"PETROL_%d",
		time.Now().UnixNano(),
	)

	transmissionCode := fmt.Sprintf(
		"AUTO_%d",
		time.Now().UnixNano(),
	)

	description := "Test description"
	powerHp := "249"

	createRequest := api_models.CreateCarRequest{
		Name: "Toyota Camry",

		BodyType: &api_models.CodeReference{
			Code: &bodyCode,
		},

		Brand: &api_models.NameReference{
			Name: &brandName,
		},

		Model: &api_models.CarModel{
			Name: &modelName,
		},

		Color: &api_models.CarColor{
			Code: &colorCode,
			Hex:  &colorHex,
		},

		DriveType: &api_models.CodeReference{
			Code: &driveCode,
		},

		FuelType: &api_models.CodeReference{
			Code: &fuelCode,
		},

		TransmissionType: &api_models.CodeReference{
			Code: &transmissionCode,
		},

		Description: &description,
		PowerHp:     &powerHp,
	}

	created, err := car.CreateCarService(
		ctx,
		createRequest,
	)

	if err != nil {
		t.Fatal(err)
	}

	if created == nil {
		t.Fatal("created car is nil")
	}

	if created.Id == "" {
		t.Fatal("car id is empty")
	}

	if created.Brand == nil {
		t.Fatal("brand is nil")
	}

	if created.Model == nil {
		t.Fatal("model is nil")
	}

	if created.Color == nil {
		t.Fatal("color is nil")
	}

	list, err := car.GetCarsService(ctx)

	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, item := range list {
		if item.Id == created.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("car not found in list")
	}

	updatedName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	updatedPower := "300"

	updated, err := car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Name:    &updatedName,
			PowerHp: &updatedPower,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Name != updatedName {
		t.Fatal("name was not updated")
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Name != updatedName {
		t.Fatal("incorrect name after update")
	}

	if actual.PowerHp == nil {
		t.Fatal("power_hp is nil")
	}

	if *actual.PowerHp != updatedPower {
		t.Fatal("power_hp was not updated")
	}
}

func Test_Car_Service_Create_WithNewReferences(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"LADA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"GRANTA_%d",
		time.Now().UnixNano(),
	)

	bodyCode := fmt.Sprintf(
		"SEDAN_%d",
		time.Now().UnixNano(),
	)

	driveCode := fmt.Sprintf(
		"FWD_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"PETROL_%d",
		time.Now().UnixNano(),
	)

	transmissionCode := fmt.Sprintf(
		"AUTO_%d",
		time.Now().UnixNano(),
	)

	carCreated, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Lada Granta",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},

			BodyType: &api_models.CodeReference{
				Code: &bodyCode,
			},

			DriveType: &api_models.CodeReference{
				Code: &driveCode,
			},

			FuelType: &api_models.CodeReference{
				Code: &fuelCode,
			},

			TransmissionType: &api_models.CodeReference{
				Code: &transmissionCode,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if carCreated.Brand == nil {
		t.Fatal("brand is nil")
	}

	if carCreated.Brand.Name == nil {
		t.Fatal("brand name is nil")
	}

	if *carCreated.Brand.Name != brandName {
		t.Fatal("incorrect brand")
	}

	if carCreated.Model == nil {
		t.Fatal("model is nil")
	}

	if carCreated.Model.Name != modelName {
		t.Fatal("incorrect model")
	}
}

func Test_Car_Service_Update_ChangeBrandAndModel(t *testing.T) {
	ctx := context.Background()

	firstBrand := fmt.Sprintf(
		"TOYOTA_%d",
		time.Now().UnixNano(),
	)

	firstModel := fmt.Sprintf(
		"CAMRY_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Test Car",

			Brand: &api_models.NameReference{
				Name: &firstBrand,
			},

			Model: &api_models.CarModel{
				Name: &firstModel,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	secondBrand := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	secondModel := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	_, err = car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Brand: &api_models.NameReference{
				Name: &secondBrand,
			},

			Model: &api_models.CarModel{
				Name: &secondModel,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand == nil {
		t.Fatal("brand is nil")
	}

	if actual.Brand.Name == nil {
		t.Fatal("brand name is nil")
	}

	if *actual.Brand.Name != secondBrand {
		t.Fatal("brand not updated")
	}

	if actual.Model == nil {
		t.Fatal("model is nil")
	}

	if actual.Model.Name != secondModel {
		t.Fatal("model not updated")
	}
}

func Test_Car_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := car.GetCarByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Car_Service_Delete(t *testing.T) {
	ctx := context.Background()

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Delete Car",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	err = car.DeleteCarService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err == nil {
		t.Fatal("expected not found error")
	}
}

func Test_Car_Service_Update_ClearBrand(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"TOYOTA_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Clear Brand",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	updated, err := car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Brand: &api_models.NameReference{},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Brand != nil {
		t.Fatal("brand was not cleared")
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand != nil {
		t.Fatal("brand was not cleared in database")
	}
}

func Test_Car_Service_Update_ClearModel(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Clear Model",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	updated, err := car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Model: &api_models.CarModel{},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Model != nil {
		t.Fatal("model was not cleared")
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Model != nil {
		t.Fatal("model was not cleared in database")
	}
}
func Test_Car_Service_Create_ModelInheritsBrand(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"TOYOTA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"CAMRY_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Toyota Camry",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand == nil {
		t.Fatal("car brand is nil")
	}

	if actual.Model == nil {
		t.Fatal("car model is nil")
	}

	if actual.Model.Brand == nil {
		t.Fatal("model brand is nil")
	}

	if actual.Model.Brand.Name == nil {
		t.Fatal("model brand name is nil")
	}

	if *actual.Model.Brand.Name != brandName {
		t.Fatalf(
			"expected model brand %q got %q",
			brandName,
			*actual.Model.Brand.Name,
		)
	}
}
func Test_Car_Service_Update_ClearAllReferences(t *testing.T) {
	ctx := context.Background()

	bodyCode := fmt.Sprintf(
		"BODY_%d",
		time.Now().UnixNano(),
	)

	brandName := fmt.Sprintf(
		"TOYOTA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"CAMRY_%d",
		time.Now().UnixNano(),
	)

	colorCode := fmt.Sprintf(
		"BLACK_%d",
		time.Now().UnixNano(),
	)

	driveCode := fmt.Sprintf(
		"FWD_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"PETROL_%d",
		time.Now().UnixNano(),
	)

	transmissionCode := fmt.Sprintf(
		"AUTO_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Full Car",

			BodyType: &api_models.CodeReference{
				Code: &bodyCode,
			},

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},

			Color: &api_models.CarColor{
				Code: &colorCode,
			},

			DriveType: &api_models.CodeReference{
				Code: &driveCode,
			},

			FuelType: &api_models.CodeReference{
				Code: &fuelCode,
			},

			TransmissionType: &api_models.CodeReference{
				Code: &transmissionCode,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			BodyType:         &api_models.CodeReference{},
			Brand:            &api_models.NameReference{},
			Model:            &api_models.CarModel{},
			Color:            &api_models.CarColor{},
			DriveType:        &api_models.CodeReference{},
			FuelType:         &api_models.CodeReference{},
			TransmissionType: &api_models.CodeReference{},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.BodyType != nil {
		t.Fatal("body_type was not cleared")
	}

	if actual.Brand != nil {
		t.Fatal("brand was not cleared")
	}

	if actual.Model != nil {
		t.Fatal("model was not cleared")
	}

	if actual.Color != nil {
		t.Fatal("color was not cleared")
	}

	if actual.DriveType != nil {
		t.Fatal("drive_type was not cleared")
	}

	if actual.FuelType != nil {
		t.Fatal("fuel_type was not cleared")
	}

	if actual.TransmissionType != nil {
		t.Fatal("transmission_type was not cleared")
	}
}
func Test_Car_Service_Update_NameOnly_KeepsReferences(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Original Name",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := "Updated Name"

	_, err = car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Name: &newName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Name != newName {
		t.Fatal("name not updated")
	}

	if actual.Brand == nil {
		t.Fatal("brand disappeared")
	}

	if actual.Model == nil {
		t.Fatal("model disappeared")
	}

	if actual.Brand.Name == nil {
		t.Fatal("brand name missing")
	}

	if *actual.Brand.Name != brandName {
		t.Fatal("brand changed unexpectedly")
	}

	if actual.Model.Name != modelName {
		t.Fatal("model changed unexpectedly")
	}
}

func Test_Car_Service_Create_ResponseContainsRelations(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Create Response Test",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created.Brand == nil {
		t.Fatal("brand missing in create response")
	}

	if created.Model == nil {
		t.Fatal("model missing in create response")
	}

	if created.Brand.Name == nil {
		t.Fatal("brand name missing in create response")
	}

	if *created.Brand.Name != brandName {
		t.Fatal("incorrect brand in create response")
	}

	if created.Model.Name != modelName {
		t.Fatal("incorrect model in create response")
	}
}

func Test_Car_Service_Update_ResponseContainsRelations(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Update Response Test",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	updated, err := car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Name: &newName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Name != newName {
		t.Fatal("name not updated")
	}

	if updated.Brand == nil {
		t.Fatal("brand missing in update response")
	}

	if updated.Model == nil {
		t.Fatal("model missing in update response")
	}

	if updated.Brand.Name == nil {
		t.Fatal("brand name missing in update response")
	}

	if *updated.Brand.Name != brandName {
		t.Fatal("brand changed unexpectedly")
	}

	if updated.Model.Name != modelName {
		t.Fatal("model changed unexpectedly")
	}
}

func Test_Car_Service_Create_ResponseContainsAllRelations(t *testing.T) {
	ctx := context.Background()

	bodyCode := fmt.Sprintf(
		"BODY_%d",
		time.Now().UnixNano(),
	)

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	colorCode := fmt.Sprintf(
		"BLACK_%d",
		time.Now().UnixNano(),
	)

	colorHex := "#000000"

	driveCode := fmt.Sprintf(
		"AWD_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"DIESEL_%d",
		time.Now().UnixNano(),
	)

	transmissionCode := fmt.Sprintf(
		"AUTO_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Create Response Full",

			BodyType: &api_models.CodeReference{
				Code: &bodyCode,
			},

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},

			Color: &api_models.CarColor{
				Code: &colorCode,
				Hex:  &colorHex,
			},

			DriveType: &api_models.CodeReference{
				Code: &driveCode,
			},

			FuelType: &api_models.CodeReference{
				Code: &fuelCode,
			},

			TransmissionType: &api_models.CodeReference{
				Code: &transmissionCode,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created.BodyType == nil {
		t.Fatal("body_type missing in create response")
	}

	if created.Brand == nil {
		t.Fatal("brand missing in create response")
	}

	if created.Model == nil {
		t.Fatal("model missing in create response")
	}

	if created.Color == nil {
		t.Fatal("color missing in create response")
	}

	if created.DriveType == nil {
		t.Fatal("drive_type missing in create response")
	}

	if created.FuelType == nil {
		t.Fatal("fuel_type missing in create response")
	}

	if created.TransmissionType == nil {
		t.Fatal("transmission_type missing in create response")
	}

	if created.Brand.Name == nil ||
		*created.Brand.Name != brandName {
		t.Fatal("invalid brand in create response")
	}

	if created.Model.Name != modelName {
		t.Fatal("invalid model in create response")
	}

	if created.BodyType.Code == nil ||
		*created.BodyType.Code != bodyCode {
		t.Fatal("invalid body_type in create response")
	}

	if created.Color.Code == nil ||
		*created.Color.Code != colorCode {
		t.Fatal("invalid color in create response")
	}

	if created.DriveType.Code == nil ||
		*created.DriveType.Code != driveCode {
		t.Fatal("invalid drive_type in create response")
	}

	if created.FuelType.Code == nil ||
		*created.FuelType.Code != fuelCode {
		t.Fatal("invalid fuel_type in create response")
	}

	if created.TransmissionType.Code == nil ||
		*created.TransmissionType.Code != transmissionCode {
		t.Fatal("invalid transmission_type in create response")
	}
}
func Test_Car_Service_Update_ResponseContainsAllRelations(t *testing.T) {
	ctx := context.Background()

	bodyCode := fmt.Sprintf(
		"BODY_%d",
		time.Now().UnixNano(),
	)

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	colorCode := fmt.Sprintf(
		"BLACK_%d",
		time.Now().UnixNano(),
	)

	driveCode := fmt.Sprintf(
		"AWD_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"DIESEL_%d",
		time.Now().UnixNano(),
	)

	transmissionCode := fmt.Sprintf(
		"AUTO_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Update Response Full",

			BodyType: &api_models.CodeReference{
				Code: &bodyCode,
			},

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},

			Color: &api_models.CarColor{
				Code: &colorCode,
			},

			DriveType: &api_models.CodeReference{
				Code: &driveCode,
			},

			FuelType: &api_models.CodeReference{
				Code: &fuelCode,
			},

			TransmissionType: &api_models.CodeReference{
				Code: &transmissionCode,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	updated, err := car.UpdateCarService(
		ctx,
		created.Id,
		api_models.UpdateCarRequest{
			Name: &newName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Name != newName {
		t.Fatal("name not updated")
	}

	if updated.BodyType == nil {
		t.Fatal("body_type missing after update")
	}

	if updated.Brand == nil {
		t.Fatal("brand missing after update")
	}

	if updated.Model == nil {
		t.Fatal("model missing after update")
	}

	if updated.Color == nil {
		t.Fatal("color missing after update")
	}

	if updated.DriveType == nil {
		t.Fatal("drive_type missing after update")
	}

	if updated.FuelType == nil {
		t.Fatal("fuel_type missing after update")
	}

	if updated.TransmissionType == nil {
		t.Fatal("transmission_type missing after update")
	}

	if updated.Brand.Name == nil ||
		*updated.Brand.Name != brandName {
		t.Fatal("brand changed unexpectedly")
	}

	if updated.Model.Name != modelName {
		t.Fatal("model changed unexpectedly")
	}

	if updated.BodyType.Code == nil ||
		*updated.BodyType.Code != bodyCode {
		t.Fatal("body_type changed unexpectedly")
	}

	if updated.Color.Code == nil ||
		*updated.Color.Code != colorCode {
		t.Fatal("color changed unexpectedly")
	}

	if updated.DriveType.Code == nil ||
		*updated.DriveType.Code != driveCode {
		t.Fatal("drive_type changed unexpectedly")
	}

	if updated.FuelType.Code == nil ||
		*updated.FuelType.Code != fuelCode {
		t.Fatal("fuel_type changed unexpectedly")
	}

	if updated.TransmissionType.Code == nil ||
		*updated.TransmissionType.Code != transmissionCode {
		t.Fatal("transmission_type changed unexpectedly")
	}
}
