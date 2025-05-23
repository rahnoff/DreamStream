REVOKE CREATE ON SCHEMA public FROM PUBLIC;

CREATE SCHEMA IF NOT EXISTS enrollments;

CREATE TABLE IF NOT EXISTS enrollments.employees
(
    id         uuid        PRIMARY KEY,
    create_at  timestamptz NOT NULL,
    edited_at  timestamptz NOT NULL,
    first_name text        NOT NULL,
    last_name  text        NOT NULL
);

CREATE TABLE IF NOT EXISTS enrollments.courses
(
    id        uuid        PRIMARY KEY,
    create_at timestamptz NOT NULL,
    edited_at timestamptz NOT NULL,
    title     text        NOT NULL
);

CREATE TYPE IF NOT EXISTS statuses AS ENUM
(
    'Enrolled',
    'In progress',
    'Completed'
);

CREATE TABLE IF NOT EXISTS enrollments.enrollments
(
    id          uuid        PRIMARY KEY,
    course_id   uuid        NOT NULL REFERENCES enrollments.courses(id),
    create_at   timestamptz NOT NULL,
    edited_at   timestamptz NOT NULL,
    employee_id uuid        NOT NULL REFERENCES enrollments.employees(id)
    status      statuses    NOT NULL
);
