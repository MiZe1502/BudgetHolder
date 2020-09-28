SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_goods_details(amount NUMERIC,
                                                    price MONEY,
                                                    purchase_id INTEGER,
                                                    goods_item_id INTEGER,
                                                    user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_purchase_by_id(purchase_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Purchase with id % not found. ', 'ERR01', purchase_id;
    END IF;

    IF NOT EXISTS(SELECT * FROM get_goods_item_by_id(goods_item_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Goods item with id % not found. ', 'ERR01', goods_item_id;
    END IF;

    INSERT INTO goods_details(amount,
                              price,
                              purchase_id,
                              goods_id,
                              added_by_user_id)
    VALUES (amount,
            price,
            purchase_id,
            goods_item_id,
            user_by_session_id)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;