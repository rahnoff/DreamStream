CREATE DATABASE IF NOT EXISTS dream_stream ON CLUSTER cluster_1S_2R ENGINE = Atomic;


CREATE TABLE IF NOT EXISTS dream_stream.enrollments ON CLUSTER cluster_1S_2R
(
    id            UInt8,
    course_name   String,
    employee_name String,
    status        String
) ENGINE = MergeTree()
ORDER BY id
SETTINGS index_granularity = 8192;


CREATE TABLE IF NOT EXISTS dream_stream.enrollments_replicated ON CLUSTER cluster_1S_2R
(
    id            UInt8,
    course_name   String,
    employee_name String,
    status        String
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}')
ORDER BY id
SETTINGS index_granularity = 8192;
