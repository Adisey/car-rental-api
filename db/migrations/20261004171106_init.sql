-- Create "cars" table
CREATE TABLE "public"."cars" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(), 
    "name" text NOT NULL, 
    "description" text NULL, 
    PRIMARY KEY ("id")
    );
