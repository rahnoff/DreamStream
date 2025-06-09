#!/bin/bash
#
# Emulate SQL queries in PostgreSQL


function main()
{
    declare -a courses_ids=($(cat cassandra/courses.csv | cut -d ',' -f 1 | sed '1d' | sed 's/.*/"&"/' | awk 'BEGIN {ORS=" "} {print}'))
    declare -a employees_ids=($(cat postgresql/employees.csv | cut -d ',' -f 1 | sed '1d' | sed 's/.*/"&"/' | awk 'BEGIN {ORS=" "} {print}'))
    local courses_ids_size
    local employees_ids_size
    local course_id_index
    local employee_id_index
    local course_id
    local employee_id
    local running
    courses_ids_size="$((${#courses_ids[@]} - 1))"
    employees_ids_size="$((${#employees_ids[@]} - 1))"
    running=true
    while [[ "${running}" == 'true' ]]; do
        course_id_index=$(shuf -i 0-"${courses_ids_size}" -n 1)
        employee_id_index=$(shuf -i 0-"${employees_ids_size}" -n 1)
        course_id="${courses_ids[${course_id_index}]}"
        employee_id="${employees_ids[${employee_id_index}]}"
        curl -X GET http://localhost:3000/enrollments
    done
}


main "$@"