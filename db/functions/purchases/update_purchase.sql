SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_purchase(purchase_id INTEGER,
                                           updated_total_price MONEY,
                                           updated_shop_id INTEGER,
                                           updated_date TIMESTAMP,
                                           updated_comment VARCHAR(3000),
                                           user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE updated_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_purchase_by_id(purchase_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Purchase with id % not found. ', 'ERR01', purchase_id;
    END IF;

    IF NOT EXISTS(SELECT * FROM get_shop_by_id(updated_shop_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Shop with id % not found. ', 'ERR01', updated_shop_id;
    END IF;

    UPDATE purchases
    SET total_price = updated_total_price,
        shop_id = updated_shop_id,
        date = updated_date,
        comment = updated_comment,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = purchase_id
    RETURNING id INTO updated_id;

    RETURN updated_id;
END
$BODY$
    LANGUAGE plpgsql;