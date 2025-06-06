REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS enrollments;


CREATE TABLE IF NOT EXISTS enrollments.employees
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    first_name text        NOT NULL,
    last_name  text        NOT NULL,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS enrollments.categories
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS enrollments.courses
(
    id         uuid        PRIMARY KEY,
    category   uuid        NOT NULL REFERENCES enrollments.categories(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TYPE enrollments.statuses AS ENUM
(
    'Cancelled',
    'Completed',
    'Enrolled',
    'In progress'
);


CREATE TABLE IF NOT EXISTS enrollments.enrollments
(
    id          uuid                 PRIMARY KEY,
    course_id   uuid                 NOT NULL REFERENCES enrollments.courses(id) ON DELETE CASCADE,
    created_at  timestamptz          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    employee_id uuid                 NOT NULL REFERENCES enrollments.employees(id) ON DELETE CASCADE,
    status      enrollments.statuses NOT NULL,
    CHECK (edited_at >= created_at)
);


CREATE OR REPLACE FUNCTION update_edited_at() RETURNS TRIGGER AS
$$
    BEGIN
        NEW.edited_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END;
$$
LANGUAGE 'plpgsql';


DO
$$
    DECLARE
        table_name_variable text;
    BEGIN
        FOR table_name_variable IN SELECT table_name FROM information_schema.columns WHERE column_name = 'edited_at' LOOP
            EXECUTE format('CREATE TRIGGER update_edited_at
                                BEFORE UPDATE ON enrollments.%I
                                FOR EACH ROW
                                EXECUTE PROCEDURE update_edited_at()',
                            table_name_variable,
                            table_name_variable);
        END loop;
    END
$$;


INSERT INTO enrollments.courses (id, created_at, edited_at, title) VALUES ('3b854085-d26a-4f8c-90f5-36abbf1756c0', )
