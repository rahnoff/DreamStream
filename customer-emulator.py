import ast
import contextlib
import csv
import http.client
import json
import os
import random
import time


def read_ids(file_name: str) -> list[str]:
    with open(file_name, 'rt', newline='') as csv_file:
        reader: _csv.reader = csv.reader(
            csv_file,
            dialect='unix',
            delimiter=',',
            escapechar='\\',
            quoting=csv.QUOTE_NONE)
        next(reader)
        ids: list[str] = [row[0] for row in reader]
        return ids


def customer_emulator() -> None:
    enrollments_server = os.environ['ENROLLMENTS_SERVER']
    enrollments_port = os.environ['ENROLLMENTS_PORT']
    courses_ids: list[str] = read_ids('/var/tmp/courses.csv')
    employees_ids: list[str] = read_ids('/var/tmp/employees.csv')
    headers: dict[str, str] = {'Content-Type': 'application/json'}
    http_client: http.client.HTTPConnection = http.client.HTTPConnection(
        enrollments_server,
        enrollments_port)
    with contextlib.closing(http_client) as connection:
        while True:
            course_id: str = random.choice(courses_ids)
            employee_id: str = random.choice(employees_ids)
            enrollment: dict[str, str] = {
                'course_id': course_id,
                'employee_id': employee_id}
            enrollment_json = json.dumps(enrollment)
            connection.request(
                'POST',
                '/enrollments',
                enrollment_json,
                headers)
            response: str = connection.getresponse()
            if response.status == 201:
                print('Enrolled')
                created_enrollment_str: str = response.read().decode('utf-8')
                created_enrollment_dict: dict[str, str] = ast.literal_eval(created_enrollment_str)
                created_enrollment_id: str = created_enrollment_dict['id']
                enrollment_status: dict[str, str] = {'status': 'In progress'}
                enrollment_status_json = json.dumps(enrollment_status)
                created_enrollment_id_url = f"/enrollments/{created_enrollment_id}".strip()
                time.sleep(5)
                connection.request('PUT', created_enrollment_id_url, enrollment_status_json, headers)
                response: str = connection.getresponse()
                response.read()
                if response.status == 204:
                    print('Course is in progress')
                    time.sleep(5)
                    print('Sending an attempt to a quiz')
                    time.sleep(5)
                else:
                    continue
            else:
                continue


def main() -> None:
    customer_emulator()


if __name__ == '__main__':
    main()