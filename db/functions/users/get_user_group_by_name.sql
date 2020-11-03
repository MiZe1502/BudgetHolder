SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_user_group_by_name(group_name VARCHAR)
    RETURNS TABLE(id INTEGER,
                  name VARCHAR,
                  description VARCHAR) AS
$$
    SELECT id, name, description
    FROM user_group
    WHERE is_removed = FALSE AND name = group_name;
$$
LANGUAGE sql;
