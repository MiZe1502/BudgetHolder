
SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_goods_item(goods_item_id INTEGER,
                                             user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE removed_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    UPDATE goods
    SET is_removed = TRUE,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = goods_item_id
    RETURNING id INTO removed_id;

    IF removed_id IS NOT NULL THEN
        PERFORM remove_goods_details(NULL,
                                     NULL,
                                     removed_id,
                                     user_session_uuid);
    END IF;

    RETURN removed_id;
END
$BODY$
    LANGUAGE plpgsql;

