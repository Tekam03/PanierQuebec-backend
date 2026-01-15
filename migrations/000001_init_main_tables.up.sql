CREATE TABLE store_merchants (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  url text
);

CREATE TABLE stores (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  merchant_id int NOT NULL,
  name text,
  location text
);

CREATE TABLE brands (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text
);

CREATE TABLE generic_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text,
  description text,
  image_url text
);

CREATE TABLE specific_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  generic_products_id int,
  brand_id int,
  name text,
  description text,
  image_url text
);

CREATE TABLE external_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  source text NOT NULL, -- e.g. 'store_1', 'store_2'
  external_id text NOT NULL, -- the store product code
  name text,
  description text,
  brand text,
  matched_store_product_id int,
  scraped_at timestamptz DEFAULT now(),
  UNIQUE (external_id)
);

CREATE TABLE store_products (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_id int NOT NULL,
  specific_products_id int NOT NULL,
  price numeric(10, 2),
  last_updated timestamptz DEFAULT now()
);

CREATE TABLE price_histories (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_products_id int NOT NULL,
  price numeric(10, 2),
  date_recorded timestamptz DEFAULT now()
);
