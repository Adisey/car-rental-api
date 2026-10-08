-- Modify "cars" table
ALTER TABLE "public"."cars" ADD COLUMN "color_id" uuid NULL, ADD CONSTRAINT "cars_color_id_fkey" FOREIGN KEY ("color_id") REFERENCES "public"."colors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
