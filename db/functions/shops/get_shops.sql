SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_shops(start INTEGER,
                                     count INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR,
                  url VARCHAR,
                  address VARCHAR,
                  comment VARCHAR) AS
$$
    SELECT id, name, url, address, comment
    FROM shops
    WHERE is_removed = FALSE
    ORDER BY name
    LIMIT count
    OFFSET start;
$$
LANGUAGE sql;
