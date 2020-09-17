SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_user(old_login VARCHAR(100),
                                       new_login VARCHAR(100),
                                       is_password_changed BOOLEAN,
                                       new_password VARCHAR(100),
                                       new_path_to_photo VARCHAR(200),
                                       new_name VARCHAR(100),
                                       new_surname VARCHAR(100),
                                       new_description VARCHAR(3000))
    RETURNS INTEGER AS
$BODY$
DECLARE updated_user_id INTEGER;
    DECLARE updated_user_data_id INTEGER;
    DECLARE updated_user_profile_id INTEGER;
BEGIN
    IF NOT EXISTS(SELECT * FROM is_user_exists(old_login)) THEN
        RAISE EXCEPTION 'Error code: %. Description: User with login % not found. ', 'ERR01', old_login;
    END IF;

    SELECT users.id
    FROM users
             JOIN user_profile ON users.user_profile_id = user_profile.id
    WHERE user_profile.login = old_login
    INTO updated_user_id;

    SELECT user_data_id
    FROM users
    WHERE id = updated_user_id
    INTO updated_user_data_id;

    IF old_login <> new_login OR is_password_changed THEN
        SELECT user_profile_id
        FROM users
        WHERE id = updated_user_id
        INTO updated_user_profile_id;

        IF old_login <> new_login THEN
            UPDATE user_profile
            SET login = new_login
            WHERE id = updated_user_profile_id;
        END IF;

        IF is_password_changed THEN
            UPDATE user_profile
            SET password = new_password
            WHERE id = updated_user_profile_id;
        END IF;

        UPDATE user_profile
        SET path_to_photo = new_path_to_photo
        WHERE id = updated_user_profile_id;
    END IF;

    UPDATE user_data
    SET name = new_name,
        surname = new_surname
    WHERE id = updated_user_data_id;

    UPDATE users
    SET description = new_description,
        updated_at = now(),
        last_online_at = now()
    WHERE id = updated_user_id;

    RETURN updated_user_id;
END
$BODY$
    LANGUAGE plpgsql;