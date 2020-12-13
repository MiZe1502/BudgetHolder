SET search_path TO budget;

CREATE OR REPLACE FUNCTION remove_purchase(purchase_id INTEGER,
                                           user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
DECLARE user_by_session_id INTEGER;
    DECLARE removed_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_purchase_by_id(purchase_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Purchase with id % not found. ', 'ERR01', purchase_id;
    END IF;

    UPDATE purchases
    SET is_removed = TRUE,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = purchase_id
      AND added_by_user_id = user_by_session_id
    RETURNING id INTO removed_id;

    IF removed_id IS NOT NULL THEN
        PERFORM remove_goods_details(NULL,
                                     removed_id,
                                     NULL,
                                     user_session_uuid);
    END IF;

    RETURN removed_id;
END
$BODY$
    LANGUAGE plpgsql;

