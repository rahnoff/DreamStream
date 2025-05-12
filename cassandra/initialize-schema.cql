CREATE KEYSPACE IF NOT EXISTS dream_stream WITH REPLICATION =
{ 
    'class': 'NetworkTopologyStrategy', 
    'datacenter1': 1 
}
AND DURABLE_WRITES = true;

CREATE TABLE dream_stream.courses
(
    id          UUID,
    authors     set<text>,
    created_at  timestamp,
    description text,
    language    text,
    title       text,
    PRIMARY KEY (id, title)
);
