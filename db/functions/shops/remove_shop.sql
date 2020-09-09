SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_shop(shop_id INTEGER,
                                       user_session_id INTEGER)
RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
DECLARE removed_id INTEGER;
BEGIN
    SELECT user_id INTO user_by_session_id
    FROM budget.sessions
    WHERE id = user_session_id AND is_active = TRUE;

    IF user_by_session_id IS NULL THEN
        RAISE EXCEPTION 'Error code: %. Description: No users found for session: %. ', 'ERR01', user_session_id;
    END IF;

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