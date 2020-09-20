SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_goods_categories(start INTEGER,
                                                count INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR(150),
                  comment VARCHAR(3000),
                  parent_id INTEGER) AS
$$
    SELECT id, name, comment, parent_id
    FROM goods_categories
    WHERE is_removed = FALSE
    ORDER BY name
    LIMIT count
    OFFSET start;
$$
    LANGUAGE sql;