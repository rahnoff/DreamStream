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

INSERT INTO dream_stream.courses (id, authors, created_at, description, language, title, edited_at)
    VALUES (uuid(), {'Adfsdf Gdgfdg', 'Ffgfdg Dfdf'}, toTimestamp(now()), 'Dfggfdhfghfghgfhfgh fghfghfg fghfghfghfg hgfhgfhfgjghjkdfgdfg', 'EN', 'Hofofdg', toTimestamp(now()));

-- To create a CSV from a table
sudo bin/cqlsh localhost 9042 --execute "COPY dream_stream.courses TO 'keyspaces_sample_table.csv' WITH HEADER=true"

-- To copy from a CSV
USE dream_stream;
COPY courses FROM './keyspaces_sample_table.csv' WITH HEADER=true;
