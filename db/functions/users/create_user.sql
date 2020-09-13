SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_user_data(user_name VARCHAR(100),
                                            user_surname VARCHAR(100))
    RETURNS INTEGER AS
$BODY$
    DECLARE new_id INTEGER;
BEGIN
    INSERT INTO user_data(name, surname)
    VALUES (user_name, user_surname)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;

-- ////////////////////////////////////////////////////////////////////////

SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_user_profile(user_login VARCHAR(100),
                                               user_password VARCHAR(1000),
                                               path VARCHAR(200))
    RETURNS INTEGER AS
$BODY$
DECLARE new_id INTEGER;
BEGIN
    IF EXISTS(SELECT id
              FROM user_profile
              WHERE login = user_login) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % already exists. ', 'ERR02', user_login;
    END IF;

    INSERT INTO user_profile(login, password, path_to_photo)
    VALUES (user_login, user_password, path)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;

-- ////////////////////////////////////////////////////////////////////////

SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_user(user_login VARCHAR(100),
                                       user_password VARCHAR(100),
                                       user_name VARCHAR(100),
                                       user_surname VARCHAR(100),
                                       path_to_photo VARCHAR(200),
                                       user_description VARCHAR(3000),
                                       group_id INTEGER)
    RETURNS INTEGER AS
$BODY$
    DECLARE new_user_data_id INTEGER;
    DECLARE new_user_profile_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT *
    FROM create_user_data(user_name,
                          user_surname)
    INTO new_user_data_id;

    SELECT *
    FROM create_user_profile(user_login,
                             user_password,
                             path_to_photo)
    INTO new_user_profile_id;

    INSERT INTO users(user_data_id,
                      user_profile_id,
                      created_at,
                      description,
                      user_group_id)
    VALUES (new_user_data_id,
            new_user_profile_id,
            now(),
            user_description,
            group_id)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;