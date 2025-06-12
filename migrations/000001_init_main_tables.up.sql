CREATE TABLE store_merchants (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name varchar,
  url varchar
);

CREATE TABLE stores (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  merchant_id int NOT NULL,
  name varchar,
  location varchar
);

CREATE TABLE brands (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name varchar
);

CREATE TABLE generic_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name varchar,
  description varchar,
  image_url varchar
);

CREATE TABLE specific_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  generic_products_id int NOT NULL,
  brand_id int,
  name varchar,
  description varchar,
  image_url varchar
);

CREATE TABLE store_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_id int NOT NULL,
  specific_products_id int NOT NULL,
  price decimal,
  last_updated date
);

CREATE TABLE price_histories (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_products_id int NOT NULL,
  price decimal,
  date_recorded date
);
