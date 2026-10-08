-- Create "colors" table
CREATE TABLE "public"."colors" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "code" text NOT NULL, PRIMARY KEY ("id"));

CREATE UNIQUE INDEX idx_colors_code
ON colors(code);

INSERT INTO colors (code)
VALUES
('black'),
('white'),
('silver'),
('gray'),
('red'),
('blue'),
('green'),
('yellow'),
('orange'),
('brown'),
('beige'),
('gold'),
('purple');