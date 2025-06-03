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


CREATE TABLE IF NOT EXISTS enrollments.categories
(
    id         uuid        PRIMARY KEY,
    created_at timestamptz NOT NULL,
    edited_at  timestamptz NOT NULL,
    name       text        NOT NULL
);


CREATE TABLE IF NOT EXISTS enrollments.courses
(
    id         uuid        PRIMARY KEY,
    category   uuid        REFERENCES enrollments.categories(id),
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
    id          uuid                 PRIMARY KEY,
    course_id   uuid                 NOT NULL REFERENCES enrollments.courses(id),
    created_at  timestamptz          NOT NULL,
    edited_at   timestamptz          NOT NULL,
    employee_id uuid                 NOT NULL REFERENCES enrollments.employees(id),
    status      enrollments.statuses NOT NULL
);


INSERT INTO enrollments.courses (id, created_at, edited_at, title) VALUES ('3b854085-d26a-4f8c-90f5-36abbf1756c0', )