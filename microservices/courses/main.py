import cassandra.cluster
import flask


def connect_to_cassandra() -> cassandra.cluster.Cluster:
    cluster: cassandra.cluster.Cluster = cassandra.cluster.Cluster(['192.168.0.14'], port=9042)
    session = cluster.connect()
    if 'db' not in g:
    return session


def run() -> None: