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
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    length      text,
    PRIMARY KEY (id, title)
);

INSERT INTO dream_stream.courses 
(
    id,
    authors,
    category,
    created_at,
    description,
    edited_at,
    filename,
    language,
    title
)
VALUES
(
    uuid(),
    {'Adfsdf Gdgfdg', 'Ffgfdg Dfdf'},
    'IT',
    toTimestamp(now()),
    'Dfggfdhfghfghgfhfgh fghfghfg fghfghfghfg hgfhgfhfgjghjkdfgdfg',
    toTimestamp(now()),
    'fdgdfg',
    'EN'
    'Thjldfgdfg'
);

-- Create a CSV file from a table
sudo /opt/cassandra/bin/cqlsh localhost 9042 --execute "COPY dream_stream.courses TO 'dream_stream_courses.csv' WITH HEADER=true"

-- Insert to a table from a CSV file
USE dream_stream;
COPY courses FROM './dream_stream_courses.csv' WITH HEADER=true;
