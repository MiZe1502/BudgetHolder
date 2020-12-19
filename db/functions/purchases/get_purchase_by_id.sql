
SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_purchase_by_id(purchase_id INTEGER)
    RETURNS TABLE(id INTEGER,
                  total_price MONEY,
                  shop_id INTEGER,
                  date TIMESTAMP,
                  comment VARCHAR(3000)) AS
$$
    SELECT DISTINCT id, total_price, shop_id, date, comment
    FROM purchases
    WHERE is_removed = FALSE
      AND id = purchase_id
$$
    LANGUAGE sql;
