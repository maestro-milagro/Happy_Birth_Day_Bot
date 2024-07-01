CREATE TABLE IF NOT EXISTS users
(
    user_name     CHAR,
    tg_user_name CHAR NOT NULL UNIQUE PRIMARY KEY,
    birth_day DATE
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    tg_id    CHAR,
    sub_tg_id CHAR
);