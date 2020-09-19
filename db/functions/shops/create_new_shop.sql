SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_shop(name VARCHAR(3000),
                                           url VARCHAR(3000),
                                           address VARCHAR(3000),
                                           comment VARCHAR(3000),
                                           user_session_uuid UUID)
RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
DECLARE new_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    INSERT INTO shops(name, url, address, comment, added_by_user_id)
    VALUES (name, url, address, comment, user_by_session_id) RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
LANGUAGE plpgsql;