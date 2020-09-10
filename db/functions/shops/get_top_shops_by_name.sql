SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_top_shops_by_name(shop_name VARCHAR,
                                                 top INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR) AS
$$
    SELECT id, name
    FROM shops
    WHERE is_removed = FALSE AND name ~ shop_name
    LIMIT top;
$$
LANGUAGE sql;