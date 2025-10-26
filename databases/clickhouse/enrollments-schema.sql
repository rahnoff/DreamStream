CREATE DATABASE IF NOT EXISTS dream_stream ON CLUSTER default ENGINE = Atomic;

CREATE TABLE IF NOT EXISTS dream_stream.enrollments_local ON CLUSTER default
(
    id                   UInt256 CODEC(T64, ZSTD(1)),
    course_name          LowCardinality(String),
    course_category      Enum8('Cancelled' = 0, 'Completed' = 1, 'Enrolled' = 2, 'In progress' = 3),
    employee_first_name  LowCardinality(String),
    employee_second_name LowCardinality(String),
    enrolled_at          DateTime64(9) CODEC(DoubleDelta, ZSTD(1)),
    status               Enum8('Cancelled' = 0, 'Completed' = 1, 'Enrolled' = 2, 'In progress' = 3),
    finished_at          DateTime64(9) CODEC(DoubleDelta, ZSTD(1))
)
ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}')
PRIMARY KEY(status)
ORDER BY(status, enrolled_at);
