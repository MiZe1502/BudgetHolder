SET search_path TO budget;

CREATE OR REPLACE FUNCTION is_user_exists(user_login VARCHAR(100))
    RETURNS TABLE(user_id INTEGER) AS
$$
    SELECT users.id
    FROM users
    JOIN user_profile ON users.user_profile_id = user_profile.id
    WHERE user_profile.login = user_login AND
          users.is_active = TRUE
$$
    LANGUAGE sql;