SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_user_id_by_session_uuid(user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
BEGIN
    SELECT user_id INTO user_by_session_id
    FROM sessions
    WHERE uuid = user_session_uuid AND is_active=TRUE;

    IF user_by_session_id IS NULL THEN
        RAISE EXCEPTION 'Error code: %. Description: No users found for session: %. ', 'ERR01', user_session_id;
    END IF;

    RETURN user_by_session_id;
END
$BODY$
    LANGUAGE plpgsql;