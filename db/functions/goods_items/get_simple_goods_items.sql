SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_simple_goods_items()
    RETURNS TABLE(id INTEGER,
                  name VARCHAR) AS
$$
    SELECT id, name
    FROM goods
    WHERE is_removed = FALSE
    ORDER BY name
$$
    LANGUAGE sql;
