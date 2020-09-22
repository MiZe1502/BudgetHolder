
SET search_path TO budget;

CREATE OR REPLACE FUNCTION update_goods_category(category_id INTEGER,
                                                 updated_name VARCHAR(150),
                                                 updated_comment VARCHAR(3000),
                                                 updated_parent_id INTEGER,
                                                 user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE updated_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    UPDATE goods_categories
    SET name = updated_name,
        comment = updated_comment,
        parent_id = updated_parent_id,
        updated_at = now(),
        updated_by_user_id = user_by_session_id
    WHERE id = category_id
    RETURNING id INTO updated_id;

    RETURN updated_id;
END
$BODY$
    LANGUAGE plpgsql;



