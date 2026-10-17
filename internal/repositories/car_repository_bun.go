package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BunCarRepository struct{}

func NewCarRepository() *BunCarRepository {
	return &BunCarRepository{}
}

func (r *BunCarRepository) GetAll(ctx context.Context) ([]db_models.Car, error) {
	var cars []db_models.Car

	err := db.BunDB.NewSelect().
		Model(&cars).
		Relation("BodyType").
		Relation("Brand").
		Relation("Color").
		Relation("DriveType").
		Relation("FuelType").
		Relation("Model").
		Relation("Model.Brand").
		Relation("TransmissionType").
		Where("deleted_at IS NULL").
		Scan(ctx)

	return cars, err
}

func (r *BunCarRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Car, error) {
	car := new(db_models.Car)
	err := db.BunDB.NewSelect().
		Model(car).
		Relation("BodyType").
		Relation("Brand").
		Relation("Color").
		Relation("DriveType").
		Relation("FuelType").
		Relation("Model").
		Relation("Model.Brand").
		Relation("TransmissionType").
		Where("car.id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (r *BunCarRepository) CreateRepository(
	ctx context.Context,
	car *db_models.Car,
) error {
	_, err := db.BunDB.NewInsert().
		Model(car).
		Exec(ctx)
	return err
}

type CarRelationsToUpdate struct {
	BodyID             *uuid.UUID
	UpdateBody         *bool
	BrandID            *uuid.UUID
	UpdateBrand        *bool
	ColorID            *uuid.UUID
	UpdateColor        *bool
	DriveID            *uuid.UUID
	UpdateDrive        *bool
	FuelID             *uuid.UUID
	UpdateFuel         *bool
	ModelID            *uuid.UUID
	UpdateModel        *bool
	TransmissionID     *uuid.UUID
	UpdateTransmission *bool
}

func (r *BunCarRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateCarRequest,
	relations CarRelationsToUpdate,
) (*db_models.Car, error) {
	car, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Name != nil {
		car.Name = *request.Name
	}
	if request.Description != nil {
		car.Description = request.Description
	}
	if request.PowerHp != nil {
		car.PowerHp = request.PowerHp
	}
	if relations.UpdateBody != nil {
		car.BodyTypeID = relations.BodyID
	}
	if relations.UpdateBrand != nil {
		car.BrandID = relations.BrandID
	}
	if relations.UpdateColor != nil {
		car.ColorID = relations.ColorID
	}
	if relations.UpdateDrive != nil {
		car.DriveID = relations.DriveID
	}
	if relations.UpdateFuel != nil {
		car.FuelID = relations.FuelID
	}
	if relations.UpdateModel != nil {
		car.ModelID = relations.ModelID
	}
	if relations.UpdateTransmission != nil {
		car.TransmissionID = relations.TransmissionID
	}
	car.UpdatedAt = time.Now().UTC()
	_, err = db.BunDB.NewUpdate().
		Model(car).
		WherePK().
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (r *BunCarRepository) DeleteRepository(
	ctx context.Context,
	id uuid.UUID,
) error {
	car, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	car.DeletedAt = &now
	car.UpdatedAt = now
	_, err = db.BunDB.NewUpdate().
		Model(car).
		WherePK().
		Where("deleted_at IS NULL").
		Exec(ctx)
	return err
}
