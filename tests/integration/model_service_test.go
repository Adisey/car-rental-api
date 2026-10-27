//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
)

func Test_Model_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	brandCreated, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"TOYOTA_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	createRequest := api_models.CreateModelRequest{
		Name: fmt.Sprintf(
			"CAMRY_%d",
			time.Now().UnixNano(),
		),
		Brand: &api_models.NameReference{
			Id: &brandCreated.Id,
		},
	}

	created, err := model.CreateModelService(
		ctx,
		createRequest,
	)
	if err != nil {
		t.Fatal(err)
	}

	if created == nil {
		t.Fatal("model is nil")
	}

	if created.Brand == nil {
		t.Fatal("brand is nil")
	}

	models, err := model.GetModelsService(
		ctx,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, item := range models {
		if item.Id == created.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("model not found in list")
	}

	updatedName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	updated, err := model.UpdateModelService(
		ctx,
		created.Id,
		api_models.UpdateModelRequest{
			Name: &updatedName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Name != updatedName {
		t.Fatal("model was not updated")
	}

	actual, err := model.GetModelByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Name != updatedName {
		t.Fatal("incorrect name after update")
	}
}

func Test_Model_Service_Create_WithoutBrand(t *testing.T) {
	ctx := context.Background()

	created, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: fmt.Sprintf(
				"UNIVERSAL_%d",
				time.Now().UnixNano(),
			),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created.Brand != nil {
		t.Fatal("brand must be nil")
	}
}

func Test_Model_Service_Create_WithNewBrand(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"LADA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"GRANTA_%d",
		time.Now().UnixNano(),
	)

	created, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: modelName,
			Brand: &api_models.NameReference{
				Name: &brandName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created.Brand == nil {
		t.Fatal("brand is nil")
	}

	if created.Brand.Name == nil {
		t.Fatal("brand name is nil")
	}

	if *created.Brand.Name != brandName {
		t.Fatalf(
			"expected brand %s got %s",
			brandName,
			*created.Brand.Name,
		)
	}

	actual, err := model.GetModelByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand == nil {
		t.Fatal("brand not loaded")
	}

	if actual.Brand.Name == nil {
		t.Fatal("brand name not loaded")
	}

	if *actual.Brand.Name != brandName {
		t.Fatal("incorrect brand")
	}
}

func Test_Model_Service_Create_DuplicateNameSameBrand(t *testing.T) {
	ctx := context.Background()

	brandCreated, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"BRAND_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	name := fmt.Sprintf(
		"MODEL_%d",
		time.Now().UnixNano(),
	)

	first, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: name,
			Brand: &api_models.NameReference{
				Id: &brandCreated.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	second, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: name,
			Brand: &api_models.NameReference{
				Id: &brandCreated.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if first.Id != second.Id {
		t.Fatal("expected same model")
	}
}

func Test_Model_Service_Create_SameNameDifferentBrands(t *testing.T) {
	ctx := context.Background()

	brand1, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"TOYOTA_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	brand2, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"BMW_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	modelName := fmt.Sprintf(
		"CAMRY_%d",
		time.Now().UnixNano(),
	)

	first, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: modelName,
			Brand: &api_models.NameReference{
				Id: &brand1.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	second, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: modelName,
			Brand: &api_models.NameReference{
				Id: &brand2.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if first.Id == second.Id {
		t.Fatal("models must be different")
	}
}

func Test_Model_Service_Update_MoveToAnotherBrand(t *testing.T) {
	ctx := context.Background()

	brand1, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"TOYOTA_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	brand2, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"BMW_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	created, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: fmt.Sprintf(
				"MOVE_%d",
				time.Now().UnixNano(),
			),
			Brand: &api_models.NameReference{
				Id: &brand1.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = model.UpdateModelService(
		ctx,
		created.Id,
		api_models.UpdateModelRequest{
			Brand: &api_models.NameReference{
				Id: &brand2.Id,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := model.GetModelByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand == nil {
		t.Fatal("brand is nil")
	}

	if actual.Brand.Id == nil {
		t.Fatal("brand id is nil")
	}

	if *actual.Brand.Id != brand2.Id {
		t.Fatal("brand not changed")
	}
}

func Test_Model_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := model.GetModelByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
