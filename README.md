# DreamStream

E-Learning platform, based on microservices

## How to generate data
Run scripts in order:
- `python3 ./data-generation-scripts/postgresql/create-employees.py`
- `python3 ./data-generation-scripts/postgresql/create-authors.py`
- `python3 ./data-generation-scripts/postgresql/create-categories.py`
- `python3 ./data-generation-scripts/postgresql/create-courses.py`
- `python3 ./data-generation-scripts/postgresql/create-authors-courses.py`
- `python3 ./data-generation-scripts/postgresql/create-quizes.py`
- `python3 ./data-generation-scripts/postgresql/create-questions.py`
- `python3 ./data-generation-scripts/postgresql/create-answers.py`
- `python3 ./data-generation-scripts/postgresql/create-courses-simplified.py`
- `python3 ./data-generation-scripts/postgresql/create-quizes-simplified.py`
- `python3 ./data-generation-scripts/postgresql/create-questions-simplified.py`
- `python3 ./data-generation-scripts/postgresql/create-answers-simplified.py`
They create CSVs.

## How to prepare a database

Courses schema:
`psql -c "\copy courses.categories from '/var/tmp/categories.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.courses from '/var/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.employees from '/var/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.authors from '/var/tmp/authors.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.authors-courses from '/var/tmp/authors-courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.quizes from '/var/tmp/quizes.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.questions from '/var/tmp/questions.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy courses.answers from '/var/tmp/answers.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

Enrollments schema:
`psql -c "\copy enrollments.courses from '/var/tmp/courses.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

`psql -c "\copy enrollments.employees from '/var/tmp/employees.csv' with DELIMITER ',' CSV HEADER" -d dream_stream -h /var/run/postgresql -U postgres`

## How to run

`COURSES_URL="127.0.0.1:2002" POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/dream_stream?sslmode=disable" go run ./courses.go`

`ENROLLMENTS_URL="127.0.0.1:2003" POSTGRESQL_URL="postgres://postgres:postgres@localhost:5432/dream_stream?sslmode=disable" go run ./enrollments.go`

`ENROLLMENTS_PORT=3000 ENROLLMENTS_SERVER=linux-mint python3 ./customer-emulator.py`

`go run ./http-client-loop.go`
