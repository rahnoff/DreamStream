import time

from cassandra.cluster import Cluster


def customer_emulator() -> None:
    cluster: Cluster = Cluster(['127.0.0.1'])
    session: Session = cluster.connect('dream_stream')
    running: bool = True
    query: str = 'SELECT description, title FROM movies;'
    while running:
        movies: ResultSet = session.execute(query)
        for movie in movies:
            print(movie.description, movie.title)
            time.sleep(5)


def main() -> None:
    customer_emulator()


if __name__ == '__main__':
    main()
