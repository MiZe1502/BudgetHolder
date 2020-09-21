SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_simple_goods_categories()
    RETURNS TABLE(id INTEGER,
                  value VARCHAR(150)) AS
$$
    SELECT id, name as value
    FROM goods_categories
    WHERE is_removed = FALSE
    ORDER BY name;
$$
    LANGUAGE sql;