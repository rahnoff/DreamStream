REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS courses;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA courses;


CREATE TABLE IF NOT EXISTS courses.categories
(
    id         smallint    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS courses.courses
(
    id            int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_id   smallint    NOT NULL REFERENCES courses.categories(id) ON DELETE CASCADE,
    created_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name          text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS courses.employees
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email      text        NOT NULL UNIQUE,
    first_name text        NOT NULL,
    last_name  text        NOT NULL,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS courses.authors
(
    id uuid PRIMARY KEY REFERENCES courses.employees(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS courses.authors_courses
(
    author_id uuid NOT NULL REFERENCES courses.authors(id) ON DELETE CASCADE,
    course_id int  NOT NULL REFERENCES courses.courses(id) ON DELETE CASCADE,
    PRIMARY KEY (author_id, course_id)
);


CREATE TABLE IF NOT EXISTS courses.quizes
(
    id         int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    course_id  int         NOT NULL REFERENCES courses.courses(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       text        NOT NULL,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS courses.questions
(
    id         int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content    text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    quiz_id    int         NOT NULL REFERENCES courses.quizes(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS courses.answers
(
    id          int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content     text        NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_correct  bool        NOT NULL,
    question_id int         NOT NULL REFERENCES courses.questions(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE OR REPLACE FUNCTION courses.update_edited_at() RETURNS TRIGGER LANGUAGE plpgsql AS
$$
    BEGIN
        NEW.edited_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END
$$;


DO
$$
    DECLARE
        table_name_variable text;
    BEGIN
        FOR table_name_variable IN SELECT table_name FROM information_schema.columns WHERE column_name = 'edited_at' LOOP
            EXECUTE format('CREATE TRIGGER update_edited_at
                                BEFORE UPDATE ON courses.%I
                                FOR EACH ROW
                                EXECUTE PROCEDURE courses.update_edited_at()',
                           table_name_variable,
                           table_name_variable);
        END loop;
    END
$$;


CREATE MATERIALIZED VIEW IF NOT EXISTS courses.courses_m_v AS
    SELECT
        em.id AS employee_id,
        em.first_name || ' ' || em.last_name AS employee_name,
        co.name AS course_name,
        en.created_at AS enrolled_at,
        en.status AS enrollment_status
    FROM
        courses.courses AS co
    INNER JOIN
        courses.categories AS ca
        ON co.category_id = ca.id
    INNER JOIN
        courses.authors_courses AS au_co
        ON co.course_id = co.id
WITH DATA;


CREATE UNIQUE INDEX courses_m_v_i ON courses.courses_m_v(course_name, employee_id);


CREATE OR REPLACE FUNCTION courses.update_courses_m_v() RETURNS trigger LANGUAGE plpgsql AS
$$
    BEGIN
        REFRESH MATERIALIZED VIEW CONCURRENTLY courses.courses_m_v;
        RETURN NULL;
    END
$$;


CREATE TRIGGER update_courses_m_v AFTER INSERT OR UPDATE OR DELETE ON courses.courses
    FOR EACH STATEMENT EXECUTE PROCEDURE courses.update_courses_m_v();
