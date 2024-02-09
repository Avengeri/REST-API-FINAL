CREATE TABLE t_users
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(64) NOT NULL UNIQUE,
    username      VARCHAR(64) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created       TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT
    ON COLUMN t_users.email IS 'Email пользователя';
COMMENT
    ON COLUMN t_users.username IS 'Nickname пользователя';
COMMENT
    ON COLUMN t_users.password_hash IS 'Хеш пароля';
COMMENT
    ON COLUMN t_users.created IS 'Дата создания пользователя';

CREATE TABLE t_users_todo
(
    id   SERIAL PRIMARY KEY,
    age  INTEGER     NOT NULL,
    name VARCHAR(64) NOT NULL
);

COMMENT
    ON COLUMN t_users_todo.id IS 'ID';
COMMENT
    ON COLUMN t_users_todo.age IS 'Возраст';
COMMENT
    ON COLUMN t_users_todo.name IS 'Имя';

