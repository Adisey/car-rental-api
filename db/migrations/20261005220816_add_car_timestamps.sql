-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL, ADD COLUMN "deleted_at" timestamptz NULL;
