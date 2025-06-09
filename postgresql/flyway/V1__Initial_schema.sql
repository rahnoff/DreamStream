REVOKE CREATE ON SCHEMA public FROM PUBLIC;


-- Application schema

CREATE SCHEMA IF NOT EXISTS enrollments;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA enrollments;


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
    id            uuid        PRIMARY KEY,
    category_id   uuid        NOT NULL REFERENCES enrollments.categories(id) ON DELETE CASCADE,
    created_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name          text        NOT NULL UNIQUE,
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


-- Auxiliary stuff

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


CREATE MATERIALIZED VIEW IF NOT EXISTS enrollments.enrollments_m_view AS
    SELECT
        em.first_name || em.last_name AS employee_name,
        co.name AS course_name,
        en.created_at AS enrolled_at,
        en.status AS enrollment_status
    FROM
        enrollments.enrollments AS en
    INNER JOIN
        enrollments.employees AS em
        ON en.employee_id = em.id
    INNER JOIN
        enrollments.courses AS co
        ON en.course_id = co.id
WITH DATA;


INSERT INTO enrollments.categories (id, created_at, edited_at, name) VALUES ('c2dfe11c-e405-468a-932d-2ef78195e7f3', '2025-06-08 00:36:42.781959+03:00', '2025-06-08 00:36:42.781959+03:00', 'Finance'),
                                                                            ('586e78e9-6b23-443f-9bf3-28fb09ec8b7a', '2025-06-08 00:36:42.781959+03:00', '2025-06-08 00:36:42.781959+03:00', 'IT'),
                                                                            ('62c2aa58-e941-4b85-a982-2e01670c3207', '2025-06-08 00:36:42.781959+03:00', '2025-06-08 00:36:42.781959+03:00', 'Marketing');
