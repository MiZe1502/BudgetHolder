SET search_path TO budget;

CREATE OR REPLACE FUNCTION count_goods_items()
    RETURNS INTEGER AS
$BODY$
DECLARE total INTEGER;
BEGIN
    SELECT COUNT(*) INTO total
    FROM goods
    WHERE is_removed = FALSE;

    RETURN total;
END
$BODY$
    LANGUAGE plpgsql;