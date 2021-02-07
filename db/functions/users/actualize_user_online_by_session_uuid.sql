SET search_path TO budget;

CREATE OR REPLACE FUNCTION actualize_user_online_by_session_uuid(user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
BEGIN
    SELECT user_id INTO user_by_session_id
    FROM sessions
    WHERE uuid = user_session_uuid AND is_active=TRUE;

    UPDATE users
    SET last_online_at = now()
    WHERE id = user_by_session_id;

    RETURN user_by_session_id;
END
$BODY$
    LANGUAGE plpgsql;