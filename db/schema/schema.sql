-- atlas:pos cars[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/car.go:10-19
-- atlas:pos colors[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/color.go:8-13

CREATE TABLE "cars" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "description" text, "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, "deleted_at" timestamptz, PRIMARY KEY ("id"));
CREATE TABLE "colors" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"));
