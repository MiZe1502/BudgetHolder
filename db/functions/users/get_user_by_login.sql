SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_user_by_login(user_login VARCHAR(100))
    RETURNS TABLE(id INTEGER,
                  login VARCHAR(100),
                  password VARCHAR(100)) AS
$$
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(user_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', user_login;
    END IF;

    RETURN query
        SELECT DISTINCT users.id,
                        user_profile.login,
                        user_profile.password
        FROM users
        JOIN user_profile ON users.user_profile_id = user_profile.id
        WHERE user_profile.login = user_login AND
              users.is_active = TRUE;
END;
$$
    LANGUAGE plpgsql;
