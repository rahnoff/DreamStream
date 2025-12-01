CREATE DATABASE IF NOT EXISTS dream_stream ON CLUSTER default ENGINE = Atomic;


CREATE TABLE IF NOT EXISTS dream_stream.enrollments_local ON CLUSTER default
(
    completed_at         DateTime64(9) CODEC(DoubleDelta, ZSTD(1))
    course_name          LowCardinality(String),
    course_category      LowCardinality(String),
    employee_name        LowCardinality(String),
    enrolled_at          DateTime64(9) CODEC(DoubleDelta, ZSTD(1)),
    status               Enum8('Cancelled' = 0, 'Completed' = 1, 'Enrolled' = 2, 'In progress' = 3)
)
ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}')
ORDER BY (status, enrolled_at)
SETTINGS index_granularity = 8192;


CREATE TABLE IF NOT EXISTS dream_stream.enrollments_distributed ON CLUSTER default AS dream_stream.enrollments_local ENGINE = Distributed('default', 'dream_stream', 'enrollments_local', rand());
