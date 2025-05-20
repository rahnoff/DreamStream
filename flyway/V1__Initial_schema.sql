CREATE TABLE IF NOT EXISTS dream_stream.employees
(
    id         uuid        PRIMARY KEY,
    create_at  timestamptz NOT NULL,
    edited_at  timestamptz NOT NULL,
    first_name text        NOT NULL,
    last_name  text        NOT NULL
);

CREATE TABLE IF NOT EXISTS dream_stream.courses
(
    id        uuid        PRIMARY KEY,
    create_at timestamptz NOT NULL,
    edited_at timestamptz NOT NULL,
    title     text        NOT NULL
);

