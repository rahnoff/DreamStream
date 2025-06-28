import flask
import psycopg_pool


def connect_to_postgresql() -> psycopg_pool.pool.ConnectionPool:
    if 'db' not in flask.g:
        pool: psycopg_pool.pool.ConnectionPool = psycopg_pool.ConnectionPool(conninfo='host=linux-mint port=5432 dbname=dream_stream user=postgres password=postgres', open=True)
        flask.g.db: psycopg_pool.pool.ConnectionPool = pool
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


@enrollments.route('/enrollments/<id>', methods=['GET'])
def get_enrollment_by_id() -> list[str]:
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        enrollments: list[str] = [str(record) for record in connection.execute('SELECT course_id, employee_id, employee_name, enrollment_status FROM enrollments.enrollments WHERE id = \'%s\'')]
    return enrollments


@enrollments.route('/enrollments', methods=['POST'])
def create_enrollment() -> list[str]:
    # enrollment_id: str = flask.request.json['enrollment_id']
    # employee_id: str = flask.request.json['employee_id']
    course_id: str = flask.request.json['course_id']
    # created_at: str = flask.request.json['created_at']
    # edited_at: str = flask.request.json['edited_at']
    employee_id: str = flask.request.json['employee_id']
    # status: str = flask.request.json['status']
    # create_enrollment_query: str = 'INSERT INTO enrollments.enrollments (id, course_id, created_at, edited_at, employee_id, status) VALUES (%s, %s, %s, %s, %s, %s)'
    create_enrollment_query: str = 'CALL enroll();'
    # enrollment: tuple[str] = (enrollment_id, course_id, created_at, edited_at, employee_id, status,)
    enrollment: tuple[str] = (course_id, employee_id, status,)
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        connection.execute(create_enrollment_query, enrollment)
    return [enrollment_id, course_id, created_at, edited_at, employee_id, status]


@enrollments.route('/enrollments/<id>', methods=['PUT'])
def update_enrollment() -> list[str]:
    enrollment_id: str = flask.request.json['enrollment_id']
    enrollment_status: str = flask.request.json['enrollment_status']
    # update_enrollment_query: str = 'UPDATE enrollments.enrollments SET status = %s WHERE id = %s;'
    update_enrollment_query: str = 'CALL update_enrollemnt_status();'
    enrollment: tuple[str] = (enrollment_status, enrollment_id,)
    postgresql_connection: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()
    with postgresql_connection.connection() as connection:
        connection.execute(update_enrollment_query, enrollment)
    return [enrollment_id, enrollment_status]


if __name__ == '__main__':
    enrollments.run(debug=True, host='0.0.0.0', port=3000)
