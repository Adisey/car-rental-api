-- atlas:pos brands[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/brand.go:8-13
-- atlas:pos colors[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/color.go:8-13
-- atlas:pos cars[type=table] /Users/andreysmirnov/MyProjects/car-rental/car-rental-api/internal/db_models/car.go:10-23

CREATE TABLE "brands" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_brands_name" UNIQUE ("name"));
CREATE TABLE "colors" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_colors_code" UNIQUE ("code"));
CREATE TABLE "cars" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "description" text, "color_id" uuid, "brand_id" uuid, "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, "deleted_at" timestamptz, PRIMARY KEY ("id"), FOREIGN KEY ("brand_id") REFERENCES "brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, FOREIGN KEY ("color_id") REFERENCES "colors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
