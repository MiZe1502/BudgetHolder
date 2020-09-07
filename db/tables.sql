-- USERS

CREATE TABLE IF NOT EXISTS user_data (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(100),
    surname VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS user_profile (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    login VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    path_to_photo VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS user_group (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(3000),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    is_removed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    user_data_id INTEGER NOT NULL,
    user_profile_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    last_online_at TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    description VARCHAR(3000),
    user_group_id INTEGER NOT NULL,

    FOREIGN KEY (user_data_id) REFERENCES user_data (id),
    FOREIGN KEY (user_profile_id) REFERENCES user_profile (id),
    FOREIGN KEY (user_group_id) REFERENCES user_group (id)
);

ALTER TABLE user_group ADD COLUMN added_by_user_id INTEGER;

ALTER TABLE user_group
ADD FOREIGN KEY (added_by_user_id)
REFERENCES users(id);

-- SESSIONS

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    uuid uuid UNIQUE NOT NULL,
    user_id INTEGER NOT NULL,
    started_at TIMESTAMP NOT NULL,
    ip_address VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- NOTIFICATIONS

CREATE TABLE IF NOT EXISTS severity_type (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS notifications (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    message_text VARCHAR(3000),
    title VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    is_removed BOOLEAN DEFAULT FALSE,
    severity_type_id INTEGER NOT NULL,

    FOREIGN KEY (severity_type_id) REFERENCES severity_type (id)
);

CREATE TABLE IF NOT EXISTS notifications_queue (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    user_group_id INTEGER,
    notification_id INTEGER NOT NULL,

    FOREIGN KEY (user_group_id) REFERENCES user_group (id),
    FOREIGN KEY (notification_id) REFERENCES notifications (id)
);

-- INCOME

CREATE TABLE IF NOT EXISTS income_type (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(200) UNIQUE NOT NULL,
    description VARCHAR (3000),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    is_removed BOOLEAN DEFAULT FALSE,
    added_by_user_id INTEGER,

    FOREIGN KEY (added_by_user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS income (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(250) NOT NULL,
    description VARCHAR (3000),
    date DATE NOT NULL,
    income_type_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    is_removed BOOLEAN DEFAULT FALSE,
    added_by_user_id INTEGER NOT NULL,

    FOREIGN KEY (added_by_user_id) REFERENCES users (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (income_type_id) REFERENCES income_type (id)
);

-- SHOPS

CREATE TABLE IF NOT EXISTS shops (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(300) NOT NULL,
    url VARCHAR(3000),
    address VARCHAR(3000),
    comment VARCHAR(3000),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    is_removed BOOLEAN DEFAULT FALSE,
    added_by_user_id INTEGER NOT NULL,

    FOREIGN KEY (added_by_user_id) REFERENCES users (id)
);

-- PURCHASES

CREATE TABLE IF NOT EXISTS purchases (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    total_price MONEY NOT NULL,
    shop_id INTEGER NOT NULL,
    date DATE NOT NULL,
    comment VARCHAR(3000),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    is_removed BOOLEAN DEFAULT FALSE,
    added_by_user_id INTEGER NOT NULL,

    FOREIGN KEY (added_by_user_id) REFERENCES users (id),
    FOREIGN KEY (shop_id) REFERENCES shops (id)
);