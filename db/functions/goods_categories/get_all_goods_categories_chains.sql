SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_all_goods_categories_chains()
    RETURNS TABLE(id INTEGER,
                  name VARCHAR(150),
                  parent_id INTEGER,
                  comment VARCHAR(3000)) AS
$$
    WITH RECURSIVE
        starting (id, name, parent_id, comment) AS
            (
                SELECT gc.id, gc.name, gc.parent_id, gc.comment
                FROM budget.goods_categories AS gc
                WHERE gc.parent_id IS NULL AND NOT gc.is_removed
            ),
        descendants (id, name, parent_id, comment) AS
            (
                SELECT s.id, s.name, s.parent_id, s.comment
                FROM starting AS s
                UNION ALL
                SELECT gc.id, gc.name, gc.parent_id, gc.comment
                FROM budget.goods_categories AS gc JOIN descendants AS d ON gc.parent_id = d.id
                WHERE NOT gc.is_removed
            ),
        ancestors (id, name, parent_id, comment) AS
            (
                SELECT gc.id, gc.name, gc.parent_id, gc.comment
                FROM budget.goods_categories AS gc
                WHERE gc.id IN (SELECT parent_id FROM starting) AND NOT gc.is_removed
                UNION ALL
                SELECT gc.id, gc.name, gc.parent_id, gc.comment
                FROM budget.goods_categories AS gc JOIN ancestors AS a ON gc.id = a.parent_id
                WHERE NOT gc.is_removed
            )
        TABLE ancestors
    UNION ALL
    TABLE descendants;
$$
    LANGUAGE sql;