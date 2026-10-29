-- Create "customers" table
CREATE TABLE "public"."customers" ("id" uuid NOT NULL, "first_name" text NOT NULL, "last_name" text NULL, "phone" text NULL, "email" text NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
