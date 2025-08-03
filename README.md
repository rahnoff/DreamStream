# DreamStream

To extract employees' ids from a CSV, copy this column in Excel, paste it in a file, then run these commands:
1. `sed "s/.*/'&'/" ids_as_a_column_without_quotes.txt 1>ids_as_a_column_with_quotes.txt`
2. `sed ':a; N; $!ba; s/\n/,/g' ids_as_a_column_with_quotes.txt 1>customer-emulator.py`.
After that wrap a resulting line with () in Vim.


## How to run

`psql -c "\copy enrollments.courses from '/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy enrollments.employees from '/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`FLASK_PORT=3000 POSTGRESQL_SERVER=linux-mint POSTGRESQL_PORT=5432 POSTGRESQL_DATABASE_NAME=dream_stream POSTGRESQL_USER=postgres POSTGRESQL_PASSWORD=postgres python3 microservices/enrollments/flask-implementation/main.py`

`ENROLLMENTS_PORT=3000 ENROLLMENTS_SERVER=linux-mint python3 customer-emulator.py`
