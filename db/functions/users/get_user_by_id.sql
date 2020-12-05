SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_user_by_id(user_id INTEGER)
    RETURNS VARCHAR(100) AS
$BODY$
DECLARE user_login VARCHAR;
BEGIN
    SELECT DISTINCT user_profile.login
    FROM users
             JOIN user_profile ON users.user_profile_id = user_profile.id
    WHERE users.id = user_id AND
            users.is_active = TRUE
    INTO user_login;

    RETURN user_login;
END;
$BODY$
    LANGUAGE plpgsql;
