SET search_path TO budget;

CREATE OR REPLACE FUNCTION count_purchases()
    RETURNS INTEGER AS
$BODY$
DECLARE total INTEGER;
BEGIN
    SELECT COUNT(*) INTO total
    FROM purchases
    WHERE is_removed = FALSE;

    RETURN total;
END
$BODY$
    LANGUAGE plpgsql;