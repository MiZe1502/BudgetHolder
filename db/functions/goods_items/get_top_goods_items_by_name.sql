SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_top_goods_items_by_name(goods_item_name VARCHAR,
                                                       top INTEGER)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR) AS
$$
    SELECT id, name
    FROM goods
    WHERE is_removed = FALSE AND name ~ goods_item_name
    ORDER BY name
    LIMIT top;
$$
    LANGUAGE sql;
