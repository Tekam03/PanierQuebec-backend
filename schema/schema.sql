CREATE TABLE "store_merchants" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "url" varchar
);

CREATE TABLE "stores" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "merchant_id" int NOT NULL,
  "name" varchar,
  "location" varchar
);

CREATE TABLE "brands" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "generic_products" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "image_url" varchar
);

CREATE TABLE "specific_products" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "generic_products_id" int NOT NULL,
  "brand_id" int,
  "name" varchar,
  "description" varchar,
  "image_url" varchar
);

CREATE TABLE "store_products" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "store_id" int NOT NULL,
  "specific_products_id" int NOT NULL,
  "price" decimal,
  "last_updated" date
);

CREATE TABLE "price_histories" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "store_products_id" int NOT NULL,
  "price" decimal,
  "date_recorded" date
);

ALTER TABLE "stores" ADD FOREIGN KEY ("merchant_id") REFERENCES "store_merchants" ("id");

ALTER TABLE "specific_products" ADD FOREIGN KEY ("generic_products_id") REFERENCES "generic_products" ("id");

ALTER TABLE "store_products" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "store_products" ADD FOREIGN KEY ("specific_products_id") REFERENCES "specific_products" ("id");

ALTER TABLE "specific_products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "price_histories" ADD FOREIGN KEY ("store_products_id") REFERENCES "store_products" ("id");
