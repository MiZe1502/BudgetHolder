SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_goods_details(goods_details_id INTEGER,
                                                purchase_id_to_remove INTEGER,
                                                goods_item_id_to_remove INTEGER,
                                                user_session_uuid UUID)
    RETURNS VOID AS
$BODY$
DECLARE user_by_session_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF purchase_id_to_remove IS NOT NULL AND NOT EXISTS(SELECT * FROM get_purchase_by_id(purchase_id_to_remove)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Purchase with id % not found. ', 'ERR01', purchase_id_to_remove;
    END IF;

    IF goods_item_id_to_remove IS NOT NULL AND NOT EXISTS(SELECT * FROM get_goods_item_by_id(goods_item_id_to_remove)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Goods item with id % not found. ', 'ERR01', goods_item_id_to_remove;
    END IF;

    UPDATE goods_details
    SET is_removed = TRUE,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE is_removed = FALSE AND
        (purchase_id = purchase_id_to_remove OR
         goods_id = goods_item_id_to_remove OR
         id = goods_details_id);
END
$BODY$
    LANGUAGE plpgsql;


