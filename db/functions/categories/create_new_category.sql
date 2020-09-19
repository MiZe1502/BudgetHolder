SET search_path TO budget;

CREATE OR REPLACE FUNCTION create_new_goods_category(category_name VARCHAR(150),
                                                     category_comment VARCHAR(3000),
                                                     parent_category_id INTEGER,
                                                     user_session_uuid UUID)
    RETURNS INTEGER AS
$BODY$
    DECLARE user_by_session_id INTEGER;
    DECLARE new_id INTEGER;
BEGIN
    SELECT * INTO user_by_session_id
    FROM get_user_id_by_session_uuid(user_session_uuid);

    INSERT INTO goods_categories(name,
                                 comment,
                                 parent_id,
                                 added_by_user_id)
    VALUES (category_name,
            category_comment,
            parent_category_id,
            user_by_session_id)
    RETURNING id INTO new_id;

    RETURN new_id;
END
$BODY$
    LANGUAGE plpgsql;