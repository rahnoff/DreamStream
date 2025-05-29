REVOKE CREATE ON SCHEMA public FROM PUBLIC;


CREATE SCHEMA IF NOT EXISTS enrollments;


CREATE TABLE IF NOT EXISTS enrollments.employees
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL,
    edited_at  timestamptz NOT NULL,
    first_name text        NOT NULL,
    last_name  text        NOT NULL
);


CREATE TABLE IF NOT EXISTS enrollments.courses
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL,
    edited_at  timestamptz NOT NULL,
    title      text        NOT NULL
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
    id          uuid        PRIMARY KEY,
    course_id   uuid        NOT NULL REFERENCES enrollments.courses(id),
    created_at  timestamptz NOT NULL,
    edited_at   timestamptz NOT NULL,
    employee_id uuid        NOT NULL REFERENCES enrollments.employees(id),
    status      statuses    NOT NULL
);
