DELETE FROM store_merchants WHERE true;
-- This will remove all entries from the store_merchants table
-- Note: This is a destructive operation and will remove all data from the table


-- Reset the sequence to start from the next value
SELECT setval('store_merchants_id_seq', 1, false);
-- This will ensure that the next insert starts from 1 again