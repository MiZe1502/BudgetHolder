SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_goods_details(goods_details_id INTEGER,
                                                updated_amount NUMERIC,
                                                updated_price MONEY,
                                                updated_purchase_id INTEGER,
                                                updated_goods_id INTEGER,
                                                user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE updated_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_purchase_by_id(updated_purchase_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Purchase with id % not found. ', 'ERR01', updated_purchase_id;
    END IF;

    IF NOT EXISTS(SELECT * FROM get_goods_item_by_id(updated_goods_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Goods item with id % not found. ', 'ERR01', updated_goods_id;
    END IF;

    UPDATE goods_details
    SET amount = updated_amount,
        price = updated_price,
        category_id = updated_category_id,
        goods_id = updated_goods_id,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = goods_details_id
    RETURNING id INTO updated_id;

    RETURN updated_id;
END
$BODY$
    LANGUAGE plpgsql;
