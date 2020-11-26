SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_simple_goods_categories()
    RETURNS TABLE(id INTEGER,
                  name VARCHAR(150)) AS
$$
    SELECT id, name
    FROM goods_categories
    WHERE is_removed = FALSE
    ORDER BY name;
$$
    LANGUAGE sql;