SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_full_user_info(user_login VARCHAR(100))
    RETURNS TABLE(login VARCHAR(100),
                  name VARCHAR(100),
                  surname VARCHAR(100),
                  path_to_photo VARCHAR(200),
                  description VARCHAR(3000),
                  group_name VARCHAR(100),
                  group_description VARCHAR(3000)) AS
$$
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(user_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', user_login;
    END IF;

    RETURN query
        SELECT DISTINCT user_profile.login,
                        user_data.name,
                        user_data.surname,
                        user_profile.path_to_photo,
                        users.description,
                        user_group.name,
                        user_group.description
        FROM users
        JOIN user_profile ON users.user_profile_id = user_profile.id
        JOIN user_data ON users.user_data_id = user_data.id
        JOIN user_group ON users.user_group_id = user_group.id
        WHERE user_profile.login = user_login AND
              users.is_active = TRUE;
END;
$$
    LANGUAGE plpgsql;
