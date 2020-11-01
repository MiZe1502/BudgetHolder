SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_user_group_id_by_user_id(user_id INTEGER)
    RETURNS INTEGER AS
$BODY$
DECLARE found_user_group_id INTEGER;
BEGIN
    SELECT user_group_id INTO found_user_group_id
    FROM budget.users
    WHERE user_profile_id = user_id AND is_active=TRUE;

    IF found_user_group_id IS NULL THEN
        RAISE EXCEPTION 'Error code: %. Description: No user groups found for user: %. ', 'ERR01', user_id;
    END IF;

    RETURN found_user_group_id;
END
$BODY$
LANGUAGE plpgsql;