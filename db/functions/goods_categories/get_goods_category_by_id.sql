SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_goods_category_by_id(category_id INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR(150),
                  comment VARCHAR(3000),
                  parent_id INTEGER) AS
$$
    SELECT DISTINCT id, name, comment, parent_id
    FROM goods_categories
    WHERE is_removed = FALSE
        AND id = category_id
$$
    LANGUAGE sql;