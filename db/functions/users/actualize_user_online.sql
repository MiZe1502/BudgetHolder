SET search_path TO budget;

CREATE OR REPLACE FUNCTION actualize_user_online(user_login VARCHAR(100))
    RETURNS INTEGER AS
$BODY$
DECLARE updated_user_id INTEGER;
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(user_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', user_login;
    END IF;

    SELECT users.id
    FROM users
             JOIN user_profile ON users.user_profile_id = user_profile.id
    WHERE user_profile.login = user_login
    INTO updated_user_id;

    UPDATE users
    SET last_online_at = now()
    WHERE id = updated_user_id;

    RETURN updated_user_id;
END
$BODY$
    LANGUAGE plpgsql;