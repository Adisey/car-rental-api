-- atlas:pos cars[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/car.go:8-14

CREATE TABLE "cars" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "description" text, PRIMARY KEY ("id"));
