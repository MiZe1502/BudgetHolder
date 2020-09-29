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