SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_shop(shop_id INTEGER,
                                       updated_name VARCHAR(3000),
                                       updated_url VARCHAR(3000),
                                       updated_address VARCHAR(3000),
                                       updated_comment VARCHAR(3000),
                                       user_session_uuid UUID)
RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
DECLARE updated_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    UPDATE shops
    SET name = updated_name,
        url = updated_url,
        address = updated_address,
        comment = updated_comment,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = shop_id
    RETURNING id INTO updated_id;

    RETURN updated_id;
END
$BODY$
LANGUAGE plpgsql;