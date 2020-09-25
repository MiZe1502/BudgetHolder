SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_single_goods_item(goods_item_id INTEGER,
                                                    updated_name VARCHAR(3000),
                                                    updated_category_id INTEGER,
                                                    updated_comment VARCHAR(3000),
                                                    updated_bar_code VARCHAR(100),
                                                    user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE updated_id INTEGER;
    BEGIN
        SELECT * INTO user_by_session_id
        FROM get_user_id_by_session_uuid(user_session_uuid);

        IF NOT EXISTS(SELECT * FROM get_goods_category_by_id(updated_category_id)) THEN
            RAISE EXCEPTION 'Error code: %. Description: Category with id % not found. ', 'ERR01', updated_category_id;
        END IF;

        UPDATE goods
        SET name = updated_name,
            category_id = updated_category_id,
            bar_code = updated_bar_code,
            comment = updated_comment,
            updated_at = now(),
            updated_by_user_id = user_by_session_id
        WHERE id = goods_item_id
        RETURNING id INTO updated_id;

        RETURN updated_id;
    END
$BODY$
    LANGUAGE plpgsql;
