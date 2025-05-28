from flask import Flask
from psycopg_pool import ConnectionPool


def main() -> None:
    run()


def connect_to_postgresql() -> ConnectionPool:
    pool: ConnectionPool = ConnectionPool(conninfo='host=linux-mint port=5432 dbname=dream_stream user=postgres password=postgres',
                                          open=True)
    return pool


def run() -> None:
    enrollments: Flask = Flask(__name__)

    pool: ConnectionPool = connect_to_postgresql()


    @enrollments.route('/', methods=['GET'])
    def index() -> str:
        index: str = 'Default page'
        return index


    @enrollments.route('/enrollments', methods=['GET'])
    def get_enrollments():
        with pool.connection() as connection:
            test_str = [str(record) for record in connection.execute('SELECT employee_id, course_id FROM enrollments.enrollments;')]
        return test_str


    @enrollments.route('/enrollment', methods=['POST'])
    def create_enrollment():
        with pool.connection() as connection:


    enrollments.run(host='192.168.0.105', port=3000)


if __name__ == '__main__':
    main()
