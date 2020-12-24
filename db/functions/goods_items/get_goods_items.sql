SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_goods_items(start INTEGER,
                                           count INTEGER,
                                           user_id INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR(150),
                  comment VARCHAR(3000),
                  bar_code VARCHAR(100),
                  category_id INTEGER) AS
$$
    SELECT id, name, comment, bar_code, category_id
    FROM goods
    WHERE is_removed = FALSE
      AND added_by_user_id = user_id
    ORDER BY name
    LIMIT count
    OFFSET start;
$$
    LANGUAGE sql;
