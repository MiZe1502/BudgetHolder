SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_shop(shop_id INTEGER,
                                       user_session_uuid UUID)
RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
DECLARE removed_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    UPDATE shops
    SET is_removed = TRUE,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = shop_id
    RETURNING id INTO removed_id;

    RETURN removed_id;
END
$BODY$
LANGUAGE plpgsql;