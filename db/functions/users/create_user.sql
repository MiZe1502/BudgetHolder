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