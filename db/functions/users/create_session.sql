SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_session(user_login VARCHAR(100),
                                          user_ip_address VARCHAR(50),
                                          session_uuid UUID)
    RETURNS UUID AS
$BODY$
DECLARE user_id INTEGER;
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(user_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', user_login;
    END IF;

    SELECT id
    FROM get_user_by_login(user_login)
    INTO user_id;

    INSERT INTO sessions(uuid,
                         user_id,
                         started_at,
                         ip_address)
    VALUES (session_uuid,
            user_id,
            now(),
            user_ip_address);

    RETURN session_uuid;
END
$BODY$
    LANGUAGE plpgsql;