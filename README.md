# DreamStream

E-Learning platform, based on microservices

## How to prepare a database

`psql -c "\copy enrollments.courses from '/var/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy enrollments.employees from '/var/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`


## How to run

`COURSES_URL="127.0.0.1:2002" POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/dream_stream?sslmode=disable" go run ./courses.go`

`ENROLLMENTS_URL="127.0.0.1:2003" POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/dream_stream?sslmode=disable" go run ./enrollments.go`

`ENROLLMENTS_PORT=3000 ENROLLMENTS_SERVER=linux-mint python3 ./customer-emulator.py`