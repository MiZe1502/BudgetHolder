SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_user(user_login VARCHAR(100))
    RETURNS INTEGER AS
$BODY$
DECLARE removed_id INTEGER;
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(user_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', user_login;
    END IF;

    SELECT users.id
    FROM users
             JOIN user_profile ON users.user_profile_id = user_profile.id
    WHERE user_profile.login = user_login
    INTO removed_id;

    UPDATE users
    SET is_active = FALSE,
        updated_at = now()
    WHERE id = removed_id;

    RETURN removed_id;
END
$BODY$
    LANGUAGE plpgsql;