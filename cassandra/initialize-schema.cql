CREATE KEYSPACE IF NOT EXISTS dream_stream WITH REPLICATION =
{ 
    'class': 'NetworkTopologyStrategy', 
    'datacenter1': 1 
}
AND DURABLE_WRITES = true;

CREATE TABLE IF NOT EXISTS dream_stream.courses
(
    id          UUID,
    authors     set<text>,
    created_at  timestamp,
    description text,
    language    text,
    title       text,
    edited_at   timestamp,
    PRIMARY KEY (id, title)
);
