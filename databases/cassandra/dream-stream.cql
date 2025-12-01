CREATE KEYSPACE IF NOT EXISTS dream_stream WITH REPLICATION =
{ 
    'class': 'NetworkTopologyStrategy', 
    'datacenter_1': 1 
}
AND DURABLE_WRITES = true;


CREATE TABLE IF NOT EXISTS dream_stream.employees
(
    id          uuid,
    first_name        text,
    second_name        text,
    created_at  timestamp,
    edited_at   timestamp,
    PRIMARY KEY (id, name)
);


CREATE TABLE IF NOT EXISTS dream_stream.courses
(
    id          uuid,
    name        text,
    authors     set<text>,
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    length      text,
    PRIMARY KEY (id, name)
);


CREATE TABLE IF NOT EXISTS dream_stream.quizes
(
    id          uuid,
    name text
    course_name        text,
    authors     set<text>,
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    PRIMARY KEY (id, name)
);


CREATE TABLE IF NOT EXISTS dream_stream.questions
(
    id          uuid,
    name        text,
    authors     set<text>,
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    length      text,
    PRIMARY KEY (id, name)
);


CREATE TABLE IF NOT EXISTS dream_stream.answers
(
    id          uuid,
    name        text,
    authors     set<text>,
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    length      text,
    PRIMARY KEY (id, name)
);


CREATE TABLE IF NOT EXISTS dream_stream.reviews
(
    id          uuid,
    name        text,
    authors     set<text>,
    category    text,
    created_at  timestamp,
    description text,
    edited_at   timestamp,
    filename    text,
    language    text,
    length      text,
    PRIMARY KEY (id, name)
);