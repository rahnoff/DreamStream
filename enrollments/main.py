# from flask import Flask, request
# from psycopg_pool import ConnectionPool
import flask
import psycopg_pool


def connect_to_postgresql() -> psycopg_pool.pool.ConnectionPool:
    pool: psycopg_pool.pool.ConnectionPool = psycopg_pool.ConnectionPool(conninfo='host=linux-mint port=5432 dbname=dream_stream user=postgres password=postgres', open=True)
    return pool


def run() -> None:
    enrollments: flask.app.Flask = flask.Flask(__name__)
    pool: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()


    @enrollments.route('/', methods=['GET'])
    def index() -> str:
        index: str = 'Default page' + '\n'
        return index


    @enrollments.route('/enrollments', methods=['GET'])
    def get_enrollments() -> list[str]:
        with pool.connection() as connection:
            enrollments: list[str] = [str(record) for record in connection.execute('SELECT em.first_name, em.last_name, co.title, en.status FROM enrollments.enrollments AS en INNER JOIN enrollments.employees AS em ON en.employee_id = em.id INNER JOIN enrollments.courses AS co ON en.course_id = co.id;')]
        return enrollments


    @enrollments.route('/enrollment', methods=['POST'])
    def create_enrollment() -> list[str]:
        enrollment_id = flask.request.json['enrollment_id']
        employee_id = flask.request.json['employee_id']
        course_id = flask.request.json['course_id']
        created_at = flask.request.json['created_at']
        edited_at = flask.request.json['edited_at']
        status = flask.request.json['status']
        create_enrollment_query: str = 'INSERT INTO enrollments.enrollments (id, course_id, created_at, edited_at, employee_id, status) VALUES (%s, %s, %s, %s, %s, %s)'
        enrollment = (enrollment_id, course_id, created_at, edited_at, employee_id, status,)
        with pool.connection() as connection:
            connection.execute(create_enrollment_query, enrollment)
        return [enrollment_id, course_id, created_at, edited_at, employee_id, status]


    @enrollments.route('/enrollment', methods=['PUT'])
    def update_enrollment() -> list[str]:
        enrollment_id: str = flask.request.json['enrollment_id']
        enrollment_status: str = flask.request.json['enrollment_status']
        update_enrollment_query: str = 'UPDATE enrollments.enrollments SET status = %s WHERE id = %s;'
        enrollment: tuple[str] = (enrollment_status, enrollment_id,)
        with pool.connection() as connection:
            connection.execute(update_enrollment_query, enrollment)
        return [enrollment_status, enrollment_id]


    enrollments.run(host='0.0.0.0', port=3000)


def main() -> None:
    run()


if __name__ == '__main__':
    main()
