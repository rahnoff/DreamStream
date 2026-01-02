# DreamStream

E-Learning platform, based on microservices

## How to run

`psql -c "\copy enrollments.courses from '/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy enrollments.employees from '/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`ENROLLMENTS_URL="127.0.0.1:2000" POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" go run enrollments.go`

`ENROLLMENTS_PORT=3000 ENROLLMENTS_SERVER=linux-mint python3 customer-emulator.py`


# Place quotes to avoid parsing error

`awk -F , 'BEGIN{OFS=FS} {$4="\""$4"\""; print}' ./courses.csv.orig 1> test.csv`