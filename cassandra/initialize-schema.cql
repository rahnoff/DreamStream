CREATE KEYSPACE IF NOT EXISTS dream_stream WITH REPLICATION =
{ 
    'class': 'NetworkTopologyStrategy', 
    'datacenter1': 1 
}
AND DURABLE_WRITES = true;

CREATE TABLE IF NOT EXISTS dream_stream.courses
(
    id          UUID,
    title       text,
    authors     set<text>,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    PRIMARY KEY (id, title)
);

INSERT INTO dream_stream.courses 
(
    id,
    title,
    authors,
    created_at,
    description,
    edited_at,
    filename,
    language
)
VALUES
(
    uuid(),
    'Tjkfg dfgjkdfg',
    {'Adfsdf Gdgfdg', 'Ffgfdg Dfdf'},
    toTimestamp(now()),
    'Dfggfdhfghfghgfhfgh fghfghfg fghfghfghfg hgfhgfhfgjghjkdfgdfg',
    toTimestamp(now()),
    'fdgdfg',
    'EN'
);

-- To create a CSV from a table
sudo bin/cqlsh localhost 9042 --execute "COPY dream_stream.courses TO 'keyspaces_sample_table.csv' WITH HEADER=true"

-- To insert from a CSV
USE dream_stream;
COPY courses FROM './keyspaces_sample_table.csv' WITH HEADER=true;
