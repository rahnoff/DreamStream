REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS attempts;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA attempts;


CREATE TABLE IF NOT EXISTS attempts.categories
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.courses
(
    id            uuid        PRIMARY KEY,
    category_id   uuid        NOT NULL REFERENCES attempts.categories(id) ON DELETE CASCADE,
    created_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name          text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.employees
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email      text        NOT NULL UNIQUE,
    first_name text        NOT NULL,
    last_name  text        NOT NULL,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.authors
(
    id uuid PRIMARY KEY REFERENCES attempts.employees(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS attempts.quizes
(
    id         uuid        PRIMARY KEY,
    course_id  uuid        NOT NULL REFERENCES attempts.courses(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.questions
(
    id         uuid        PRIMARY KEY,
    content    text        NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    quiz_id    uuid        NOT NULL REFERENCES attempts.quizes(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.answers
(
    id          uuid        PRIMARY KEY,
    content     text        NOT NULL UNIQUE,
    correct     bool        NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    question_id uuid        NOT NULL REFERENCES attempts.questions(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.attempts
(
    id          uuid        PRIMARY KEY,
    content     text        NOT NULL UNIQUE,
    correct     bool        NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    question_id uuid        NOT NULL REFERENCES attempts.questions(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE OR REPLACE FUNCTION attempts.update_edited_at() RETURNS TRIGGER LANGUAGE plpgsql AS
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
                                BEFORE UPDATE ON attempts.%I
                                FOR EACH ROW
                                EXECUTE PROCEDURE attempts.update_edited_at()',
                           table_name_variable,
                           table_name_variable);
        END loop;
    END
$$;


CREATE MATERIALIZED VIEW IF NOT EXISTS enrollments.enrollments_m_v AS
    SELECT
        em.id AS employee_id,
        em.first_name || ' ' || em.last_name AS employee_name,
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


CREATE UNIQUE INDEX enrollments_m_v_index ON enrollments.enrollments_m_v(course_name, employee_id);


CREATE OR REPLACE FUNCTION enrollments.update_enrollments_m_v() RETURNS trigger LANGUAGE plpgsql AS
$$
    BEGIN
        REFRESH MATERIALIZED VIEW CONCURRENTLY enrollments.enrollments_m_v;
        RETURN NULL;
    END
$$;


CREATE TRIGGER update_enrollments_m_v AFTER INSERT OR UPDATE OR DELETE ON enrollments.enrollments
    FOR EACH STATEMENT EXECUTE PROCEDURE enrollments.update_enrollments_m_v();


CREATE OR REPLACE PROCEDURE enrollments.enroll(IN course_id_parameter uuid, IN employee_id_parameter uuid) LANGUAGE plpgsql AS
$$
    BEGIN
        INSERT INTO enrollments.enrollments
        (
            course_id,
            employee_id,
            status
        )
        VALUES
        (
            course_id_parameter,
            employee_id_parameter,
            'Enrolled'
        );
    END
$$;


CREATE OR REPLACE PROCEDURE enrollments.update_enrollment_status(IN id_parameter uuid, IN status_parameter enrollments.statuses) LANGUAGE plpgsql AS
$$
    BEGIN
        UPDATE enrollments.enrollments SET status = status_parameter WHERE id = id_parameter;
    END
$$;
