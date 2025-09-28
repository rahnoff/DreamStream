# DreamStream

## How to run

`psql -c "\copy enrollments.courses from '/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy enrollments.employees from '/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`FLASK_PORT=3000 POSTGRESQL_SERVER=linux-mint POSTGRESQL_PORT=5432 POSTGRESQL_DATABASE_NAME=dream_stream POSTGRESQL_USER=postgres POSTGRESQL_PASSWORD=postgres python3 microservices/enrollments/flask-implementation/main.py`

`ENROLLMENTS_PORT=3000 ENROLLMENTS_SERVER=linux-mint python3 customer-emulator.py`
