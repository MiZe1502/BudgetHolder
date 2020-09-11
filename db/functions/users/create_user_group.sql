SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_user_group(group_name VARCHAR(100),
                                             description VARCHAR(3000),
                                             user_session_id INTEGER)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT user_id INTO user_by_session_id
    FROM sessions
    WHERE id = user_session_id AND is_active=TRUE;

    IF user_by_session_id IS NULL THEN
        RAISE EXCEPTION 'Error code: %. Description: No users found for session: %. ', 'ERR01', user_session_id;
    END IF;

    IF EXISTS(SELECT id
              FROM user_group
              WHERE name = group_name) THEN
        RAISE EXCEPTION 'Error code: %. Description: Group with name % already exists. ', 'ERR02', group_name;
    END IF;

    INSERT INTO user_group(name, description, added_by_user_id)
    VALUES (group_name, description, user_by_session_id) RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;