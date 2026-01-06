REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS enrollments;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA enrollments;


CREATE TABLE IF NOT EXISTS enrollments.courses
(
    id            bigint         GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at    timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name          text        NOT NULL UNIQUE,
    CHECK (edited_at >= created_at)
);


CREATE TABLE IF NOT EXISTS enrollments.employees
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email      text        NOT NULL UNIQUE,
    first_name text        NOT NULL,
    last_name  text        NOT NULL,
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
    id          bigint               GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    course_id   bigint                  NOT NULL REFERENCES enrollments.courses(id) ON DELETE CASCADE,
    created_at  timestamptz          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    edited_at   timestamptz          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    employee_id uuid                 NOT NULL REFERENCES enrollments.employees(id) ON DELETE CASCADE,
    status      enrollments.statuses NOT NULL,
    CHECK (edited_at >= created_at),
    UNIQUE (course_id, employee_id)
);


CREATE OR REPLACE FUNCTION enrollments.update_edited_at() RETURNS TRIGGER LANGUAGE plpgsql AS
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
        FOR table_name_variable IN SELECT table_name FROM information_schema.columns WHERE table_catalog = 'dream_stream' AND table_schema = 'enrollments' AND column_name = 'edited_at' ORDER BY table_name ASC LOOP
            EXECUTE format('CREATE TRIGGER update_edited_at
                                BEFORE UPDATE ON enrollments.%I
                                FOR EACH ROW
                                EXECUTE PROCEDURE enrollments.update_edited_at()',
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


CREATE UNIQUE INDEX enrollments_m_v_i ON enrollments.enrollments_m_v(course_name, employee_id);


CREATE OR REPLACE FUNCTION enrollments.update_enrollments_m_v() RETURNS trigger LANGUAGE plpgsql AS
$$
    BEGIN
        REFRESH MATERIALIZED VIEW CONCURRENTLY enrollments.enrollments_m_v;
        RETURN NULL;
    END
$$;


CREATE TRIGGER update_enrollments_m_v AFTER INSERT OR UPDATE OR DELETE ON enrollments.enrollments
    FOR EACH STATEMENT EXECUTE PROCEDURE enrollments.update_enrollments_m_v();


CREATE OR REPLACE PROCEDURE enrollments.enroll(IN course_id_parameter bigint, IN employee_id_parameter uuid, OUT enrollment_id_parameter bigint) LANGUAGE plpgsql AS
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
        )
        RETURNING id INTO enrollment_id_parameter;
    END
$$;


CREATE OR REPLACE PROCEDURE enrollments.update_enrollment_status(IN id_parameter_in bigint, IN status_parameter enrollments.statuses, OUT id_parameter_out bigint) LANGUAGE plpgsql AS
$$
    BEGIN
        UPDATE enrollments.enrollments SET status = status_parameter WHERE id = id_parameter_in
            RETURNING id INTO id_parameter_out;
    END
$$;
