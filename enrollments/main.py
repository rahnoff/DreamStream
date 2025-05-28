from flask import Flask, request
from psycopg_pool import ConnectionPool


def main() -> None:
    run()


def connect_to_postgresql() -> ConnectionPool:
    pool: ConnectionPool = ConnectionPool(conninfo='host=linux-mint port=5432 dbname=dream_stream user=postgres password=postgres', open=True)
    return pool


def run() -> None:
    enrollments: Flask = Flask(__name__)

    pool: ConnectionPool = connect_to_postgresql()


    @enrollments.route('/', methods=['GET'])
    def index() -> str:
        index: str = 'Default page' + '\n'
        return index


    @enrollments.route('/enrollments', methods=['GET'])
    def get_enrollments():
        with pool.connection() as connection:
            enrollments: list[str] = [str(record) for record in connection.execute('SELECT em.first_name, em.last_name, co.title, en.status FROM enrollments.enrollments AS en INNER JOIN enrollments.employees AS em ON en.employee_id = em.id INNER JOIN enrollments.courses AS co ON en.course_id = co.id;')]
        return enrollments


    @enrollments.route('/enrollment', methods=['POST'])
    def create_enrollment():
        course_id = request.json['course_id']
        course_title = request.json['course_title']
        course_created_at = request.json['course_created_at']
        course_edited_at = request.json['course_edited_at']
        # employee_id = request.json['employee_id']
        # employee_first_name = request.json['employee_first_name']
        # employee_last_name = request.json['employee_last_name']
        # enrollment_status = request.json['enrollment_status']
        create_course = "INSERT INTO enrollments.courses (id, create_at, edited_at, title) VALUES (%s, %s, %s, %s)"
        course = (course_id, course_created_at, course_edited_at, course_title,)
        with pool.connection() as connection:
            connection.execute(create_course, course)
        return [course_id, course_title, course_created_at, course_edited_at]

    enrollments.run(host='192.168.0.105', port=3000)


if __name__ == '__main__':
    main()
