ALTER TABLE stores ADD CONSTRAINT fk_stores_merchant_id FOREIGN KEY (merchant_id) REFERENCES store_merchants (id);
ALTER TABLE specific_products ADD CONSTRAINT fk_specific_products_generic_products_id FOREIGN KEY (generic_products_id) REFERENCES generic_products (id);
ALTER TABLE store_products ADD CONSTRAINT fk_store_products_store_id FOREIGN KEY (store_id) REFERENCES stores (id);
ALTER TABLE store_products ADD CONSTRAINT fk_store_products_specific_products_id FOREIGN KEY (specific_products_id) REFERENCES specific_products (id);
ALTER TABLE specific_products ADD CONSTRAINT fk_specific_products_brand_id FOREIGN KEY (brand_id) REFERENCES brands (id);
ALTER TABLE price_histories ADD CONSTRAINT fk_price_histories_store_products_id FOREIGN KEY (store_products_id) REFERENCES store_products (id);
ALTER TABLE external_products ADD CONSTRAINT fk_external_products_matched_store_product_id FOREIGN KEY (matched_store_product_id) REFERENCES store_products (id);
