SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_shop(name VARCHAR(3000),
                                           url VARCHAR(3000),
                                           address VARCHAR(3000),
                                           comment VARCHAR(3000),
                                           user_session_id INTEGER)
RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
DECLARE new_id INTEGER;
BEGIN
    SELECT user_id INTO user_by_session_id
    FROM sessions
    WHERE id = user_session_id AND is_active=TRUE;

    IF user_by_session_id IS NULL THEN
        RAISE EXCEPTION 'Error code: %. Description: No users found for session: %. ', 'ERR01', user_session_id;
    END IF;

    INSERT INTO shops(name, url, address, comment, added_by_user_id)
    VALUES (name, url, address, comment, user_by_session_id) RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
LANGUAGE plpgsql;