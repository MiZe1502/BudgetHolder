SET search_path TO budget;

CREATE OR REPLACE FUNCTION close_session(session_uuid UUID)
    RETURNS UUID AS
$BODY$
BEGIN
    IF NOT EXISTS(SELECT *
                  FROM sessions
                  WHERE uuid = session_uuid AND
                          is_active = TRUE) THEN
        RAISE EXCEPTION 'Error code: %. Description: Session with uuid % not found. ', 'ERR01', session_uuid;
    END IF;

    UPDATE sessions
    SET is_active = FALSE
    WHERE uuid = session_uuid AND is_active = TRUE;

    RETURN session_uuid;
END
$BODY$
    LANGUAGE plpgsql;
