#!/bin/bash


courses_ids=($(cat cassandra/courses.csv | cut -d ',' -f 1 | sed '1d' | sed 's/.*/"&"/' | awk 'BEGIN {ORS=" "} {print}'))
employees_ids=($(cat postgresql/employees.csv | cut -d ',' -f 1 | sed '1d' | sed 's/.*/"&"/' | awk 'BEGIN {ORS=" "} {print}'))
while true; do course_id=$(shuf -i 0-29 -n 1); employee_id=$(shuf -i 0-9999 -n 1); echo "${course_id}" "${employee_id}"; done