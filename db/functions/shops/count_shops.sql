SET search_path TO budget;

CREATE OR REPLACE FUNCTION count_shops()
RETURNS INTEGER AS
$BODY$
    DECLARE total INTEGER;
BEGIN
    SELECT COUNT(*) INTO total
    FROM shops
    WHERE is_removed = FALSE;

    RETURN total;
END
$BODY$
    LANGUAGE plpgsql;