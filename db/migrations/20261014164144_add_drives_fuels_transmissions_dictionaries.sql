-- Create "drive_types" table
CREATE TABLE "public"."drive_types" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_drive_type_code" UNIQUE ("code"));
-- Create "fuel_types" table
CREATE TABLE "public"."fuel_types" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_fuel_type_code" UNIQUE ("code"));
-- Create "transmission_types" table
CREATE TABLE "public"."transmission_types" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_transmission_type_code" UNIQUE ("code"));
-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "drive_id" uuid NULL, ADD COLUMN "fuel_id" uuid NULL, ADD COLUMN "transmission_id" uuid NULL, ADD CONSTRAINT "cars_drive_id_fkey" FOREIGN KEY ("drive_id") REFERENCES "public"."drive_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "cars_fuel_id_fkey" FOREIGN KEY ("fuel_id") REFERENCES "public"."fuel_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "cars_transmission_id_fkey" FOREIGN KEY ("transmission_id") REFERENCES "public"."transmission_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;


INSERT INTO drive_types (code)
VALUES
('fwd'),
('rwd'),
('awd'),
('4wd'),
('6x4'),
('6x6');

INSERT INTO fuel_types (code)
VALUES
('petrol'),
('diesel'),
('hybrid'),
('plug-in-hybrid'),
('electric'),
('lpg'),
('cng'),
('hydrogen');

INSERT INTO transmission_types (code)
VALUES
('manual'),
('automatic'),
('cvt'),
('dct'),
('robot'),
('single-speed');