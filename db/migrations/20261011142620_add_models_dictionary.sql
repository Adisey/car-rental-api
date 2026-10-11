-- Create "models" table
CREATE TABLE "public"."models" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, "brand_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_models_name_brand" UNIQUE ("name", "brand_id"), CONSTRAINT "models_brand_id_fkey" FOREIGN KEY ("brand_id") REFERENCES "public"."brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "model_id" uuid NULL, ADD CONSTRAINT "cars_model_id_fkey" FOREIGN KEY ("model_id") REFERENCES "public"."models" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
