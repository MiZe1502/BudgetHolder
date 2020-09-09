SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_shop_by_id(shop_id INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR,
                  url VARCHAR,
                  address VARCHAR,
                  comment VARCHAR) AS
$$
    SELECT id, name, url, address, comment
    FROM shops
    WHERE is_removed = FALSE AND id = shop_id;
$$
LANGUAGE sql;
