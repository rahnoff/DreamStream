REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS attempts;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA attempts;


CREATE TABLE IF NOT EXISTS attempts.courses
(
    id            int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
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


CREATE TABLE IF NOT EXISTS attempts.quizes
(
    id         int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    course_id  int         NOT NULL REFERENCES attempts.courses(id) ON DELETE CASCADE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.questions
(
    id         int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    quiz_id    int         NOT NULL REFERENCES attempts.quizes(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.answers
(
    id          int         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_correct  bool        NOT NULL,
    question_id int         NOT NULL REFERENCES attempts.questions(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.attempts
(
    id          bigint      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content     jsonb       NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    employee_id uuid        NOT NULL REFERENCES attempts.employees(id) ON DELETE CASCADE,
    quiz_id     int         NOT NULL REFERENCES attempts.quizes(id) ON DELETE CASCADE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS attempts.attempts
(
    id          bigint      PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    content     jsonb       NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    employee_id uuid        NOT NULL REFERENCES attempts.employees(id) ON DELETE CASCADE,
    score       decimal(1, 1)       DEFAULT 0.0,
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


CREATE MATERIALIZED VIEW IF NOT EXISTS attempts.attempts_m_v AS
    SELECT
        em.id AS employee_id,
        em.first_name || ' ' || em.last_name AS employee_name,
        co.name AS course_name,
        at.created_at AS submitted_at
    FROM
        attempts.attempts AS at
    INNER JOIN
        attempts.employees AS em
        ON at.employee_id = em.id
    INNER JOIN
        attempts.courses AS co
        ON at.course_id = co.id
WITH DATA;


CREATE UNIQUE INDEX attempts_m_v_i ON attempts.attempts_m_v(course_name, employee_id);


CREATE OR REPLACE FUNCTION attempts.update_attempts_m_v() RETURNS trigger LANGUAGE plpgsql AS
$$
    BEGIN
        REFRESH MATERIALIZED VIEW CONCURRENTLY attempts.attempts_m_v;
        RETURN NULL;
    END
$$;


CREATE TRIGGER update_attempts_m_v AFTER INSERT OR UPDATE OR DELETE ON attempts.attempts
    FOR EACH STATEMENT EXECUTE PROCEDURE attempts.update_attempts_m_v();


CREATE OR REPLACE PROCEDURE attempts.submit(IN course_id_parameter uuid, IN employee_id_parameter uuid) LANGUAGE plpgsql AS
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


CREATE OR REPLACE PROCEDURE attempts.verify_answer(
    p_question_id INT,
    p_user_answer TEXT,
    INOUT p_is_correct BOOLEAN
)
LANGUAGE plpgsql
AS $$
DECLARE
    v_correct_answer TEXT;
BEGIN
    -- Retrieve the correct answer for the given question
    SELECT correct_answer INTO v_correct_answer
    FROM questions
    WHERE question_id = p_question_id;

    -- Check if the user's answer matches the correct answer (case-insensitive comparison)
    IF v_correct_answer IS NOT NULL AND LOWER(p_user_answer) = LOWER(v_correct_answer) THEN
        p_is_correct := TRUE;
        RAISE NOTICE 'Answer is correct for question %.', p_question_id;
    ELSE
        p_is_correct := FALSE;
        RAISE NOTICE 'Answer is incorrect for question %.', p_question_id;
    END IF;

EXCEPTION
    WHEN NO_DATA_FOUND THEN
        RAISE EXCEPTION 'Question ID % not found.', p_question_id;
END;
$$;