import os

import flask
import psycopg_pool


def connect_to_postgresql() -> psycopg_pool.pool.ConnectionPool:
    postgresql_server = os.environ['POSTGRESQL_SERVER']
    postgresql_port = os.environ['POSTGRESQL_PORT']
    postgresql_database_name = os.environ['POSTGRESQL_DATABASE_NAME']
    postgresql_user = os.environ['POSTGRESQL_USER']
    postgresql_password = os.environ['POSTGRESQL_PASSWORD']
    if 'db' not in flask.g:
        connection_info = f"host={postgresql_server} port={postgresql_port} dbname={postgresql_database_name} user={postgresql_user} password={postgresql_password}"
        flask.g.db: psycopg_pool.pool.ConnectionPool = psycopg_pool.ConnectionPool(conninfo=connection_info, open=True)
        # flask.g.db: psycopg_pool.pool.ConnectionPool = psycopg_pool.ConnectionPool(conninfo='host=linux-mint port=5432 dbname=dream_stream user=postgres password=postgres', open=True)
    return flask.g.db


def create_flask_instance() -> flask.app.Flask:
    flask_instance: flask.app.Flask = flask.Flask(__name__)
    with flask_instance.app_context():
        connect_to_postgresql()
    return flask_instance


enrollments: flask.app.Flask = create_flask_instance()


@enrollments.route('/', methods=['GET'])
def index() -> str:
    index: str = 'Default page' + '\n'
    return index


@enrollments.route('/enrollments', methods=['GET'])
def get_enrollments() -> list[str]:
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        enrollments: list[str] = [str(record) for record in connection.execute('SELECT employee_id, employee_name, course_name, enrolled_at, enrollment_status FROM enrollments.enrollments_m_v;')]
    return enrollments


@enrollments.route('/enrollments/<employee_id>', methods=['GET'])
def get_enrollments_by_employee_id(employee_id) -> list[str]:
    employee_id: tuple[str] = (employee_id,)
    get_enrollments_by_employee_id_query: str = 'SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = %s;'
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        # enrollments: list[str] = [str(record) for record in connection.execute(get_enrollments_by_employee_id_query, employee_id)]
        # enrollments: list[str] = [str(record) for record in connection.execute(get_enrollments_by_employee_id_query, employee_id).fetchall()]
        enrollments: list[str] = connection.execute(get_enrollments_by_employee_id_query, employee_id).fetchall()
    return enrollments


@enrollments.route('/enrollments', methods=['POST'])
def create_enrollment() -> list[str]:
    course_id: str = flask.request.json['course_id']
    employee_id: str = flask.request.json['employee_id']
    enrollment: tuple[str] = (course_id, employee_id,)
    create_enrollment_query: str = 'CALL enrollments.enroll(%s, %s);'
    get_created_enrollment_query: str = 'SELECT id FROM enrollments.enrollments WHERE course_id = %s AND employee_id = %s;'
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        connection.execute(create_enrollment_query, enrollment)
        # enrollment: list[str] = [str(record) for record in connection.execute(get_created_enrollment_query, enrollment)]
        enrollment_id: list[str] = [str(record) for record in connection.execute(get_created_enrollment_query, enrollment).fetchone()]
    return enrollment_id


@enrollments.route('/enrollments/<id>', methods=['PUT'])
def update_enrollment(id) -> list[str]:
    status: str = flask.request.json['status']
    enrollment: tuple[str] = (id, status,)
    id_tuple: tuple[str] = (id,)
    update_enrollment_query: str = 'CALL enrollments.update_enrollment_status(%s, %s);'
    get_updated_enrollment_query: str = 'SELECT status FROM enrollments.enrollments WHERE id = %s;'
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        connection.execute(update_enrollment_query, enrollment)
        enrollment: list[str] = [str(record) for record in connection.execute(get_updated_enrollment_query, id_tuple)]
    return enrollment


if __name__ == '__main__':
    flask_port = os.environ['FLASK_PORT']
    enrollments.run(host='0.0.0.0', port=flask_port)
