SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_simple_shops()
    RETURNS TABLE(id INTEGER,
                  name VARCHAR) AS
$$
    SELECT id, name
    FROM shops
    WHERE is_removed = FALSE
    ORDER BY name
$$
    LANGUAGE sql;
