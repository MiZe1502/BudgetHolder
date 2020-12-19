SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_purchase(total_price MONEY,
                                               shop_id INTEGER,
                                               date TIMESTAMP,
                                               comment VARCHAR(3000),
                                               user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_shop_by_id(shop_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Shop with id % not found. ', 'ERR01', shop_id;
    END IF;

    INSERT INTO purchases(total_price, shop_id, date, comment, added_by_user_id)
    VALUES (total_price, shop_id, date, comment, user_by_session_id)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;