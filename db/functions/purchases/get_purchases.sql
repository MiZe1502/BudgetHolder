SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_purchases(start INTEGER,
                                         count INTEGER,
                                         user_id INTEGER)
    RETURNS TABLE(id INTEGER,
                  total_price MONEY,
                  shop_id INTEGER,
                  date DATE,
                  comment VARCHAR(3000)) AS
$$
    SELECT id, total_price, shop_id, date, comment
    FROM purchases
    WHERE is_removed = FALSE
      AND added_by_user_id = user_id
    ORDER BY created_at DESC
    LIMIT count
    OFFSET start;
$$
    LANGUAGE sql;