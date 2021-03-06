SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_goods_item(name VARCHAR(3000),
                                                 category_id INTEGER,
                                                 comment VARCHAR(3000),
                                                 bar_code VARCHAR(100),
                                                 user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    IF NOT EXISTS(SELECT * FROM get_goods_category_by_id(category_id)) THEN
        RAISE EXCEPTION 'Error code: %. Description: Category with id % not found. ', 'ERR01', category_id;
    END IF;

    INSERT INTO goods(name, category_id, comment, bar_code, added_by_user_id)
    VALUES (name, category_id, comment, bar_code, user_by_session_id)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;