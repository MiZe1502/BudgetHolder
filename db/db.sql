-- DROP DATABASE budget_holder;

CREATE DATABASE budget_holder
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

GRANT CONNECT ON DATABASE budget_holder TO api;

REVOKE CONNECT ON DATABASE budget_holder FROM public;

