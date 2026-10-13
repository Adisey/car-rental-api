-- Modify "colors" table
ALTER TABLE "public"."colors" ADD COLUMN "hex" text NULL;
-- Create "body_types" table
CREATE TABLE "public"."body_types" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_body_type_code" UNIQUE ("code"));
-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "body_type_id" uuid NULL, ADD COLUMN "power_hp" integer NULL, ADD CONSTRAINT "cars_body_type_id_fkey" FOREIGN KEY ("body_type_id") REFERENCES "public"."body_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

INSERT INTO body_types (code)
VALUES
('sedan'),
('hatchback'),
('wagon'),
('liftback'),
('fastback'),
('coupe'),
('convertible'),
('roadster'),
('targa'),
('suv'),
('crossover'),
('pickup'),
('van'),
('minivan'),
('microvan'),
('mpv'),
('station-wagon'),
('off-road'),
('limousine'),
('double-cab-pickup');