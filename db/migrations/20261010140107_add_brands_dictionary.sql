-- Create "brands" table
CREATE TABLE "public"."brands" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_brands_name" UNIQUE ("name"));
-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "brand_id" uuid NULL, ADD CONSTRAINT "cars_brand_id_fkey" FOREIGN KEY ("brand_id") REFERENCES "public"."brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

INSERT INTO brands (name)
VALUES
('Audi'),
('BMW'),
('Citroen'),
('Cupra'),
('Dacia'),
('Fiat'),
('Ford'),
('Honda'),
('Hyundai'),
('Jeep'),
('Kia'),
('Land Rover'),
('Mazda'),
('Mercedes-Benz'),
('Mini'),
('Mitsubishi'),
('Nissan'),
('Opel'),
('Peugeot'),
('Porsche'),
('Renault'),
('Seat'),
('Skoda'),
('Suzuki'),
('Tesla'),
('Toyota'),
('Volkswagen'),
('Volvo');