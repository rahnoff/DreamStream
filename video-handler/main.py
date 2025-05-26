from flask import Flask
from psycopg_pool import ConnectionPool


microservice = Flask(__name__)


def main() -> None:
    # microservice.run(host='0.0.0.0', port=3000)
    run()


def connect_to_postgresql():
    connection_info: str = 'host=linux-mint port=5432 dbname=postgres user=postgres password=postgres'
    pool = ConnectionPool(conninfo=connection_info, open=True)
    return pool


def run() -> None:
    pool: psycopg_pool.pool.ConnectionPool = connect_to_postgresql()


    @microservice.route('/')
    def index() -> str:
        test_string: str = 'Default page'
        return test_string


    @microservice.route('/enrollments', methods = ['GET'])
    def get_enrollments():
        with pool.connection() as connection:
            test_str = [str(record) for record in connection.execute('SELECT id, test FROM test;')]
        return test_str

    
    microservice.run(host='0.0.0.0', port=3000)


if __name__ == '__main__':
    # microservice.run(host='0.0.0.0', port=3000)
    main()
