CREATE OR REPLACE PROCEDURE update_enrollment_status(enrollment_id uuid,
                                                     enrollment_status enrollments.statuses) LANGUAGE plpgsql AS
$$
    BEGIN
        UPDATE enrollments.enrollments
            SET status = enrollment_status
            WHERE id = enrollment_id;
        COMMIT;
    END;
$$;
