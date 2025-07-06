import contextlib
import csv
import http.client
import json
import random


def read_courses_ids(file_name: str) -> list[str]:
    with open(file_name, 'r', newline='') as csv_file:
        reader: _csv.reader = csv.reader(csv_file, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        courses_ids: list[str] = [row[0] for row in reader]
        return courses_ids


def read_employees_ids(file_name: str) -> list[str]:
    with open(file_name, 'r', newline='') as csv_file:
        reader: _csv.reader = csv.reader(csv_file, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        employees_ids: list[str] = [row[0] for row in reader]
        return employees_ids


def customer_emulator() -> None:
    courses_ids: list[str] = read_courses_ids('/var/tmp/courses.csv')
    employees_ids: list[str] = read_employees_ids('/var/tmp/employees.csv')
    headers = {'Content-Type': 'application/json'}
    with contextlib.closing(http.client.HTTPConnection('localhost', 3000)) as connection:
        while True:
            course_id: str = random.choice(courses_ids)
            employee_id: str = random.choice(employees_ids)
            enrollment: dict[str, str] = {'course_id': course_id, 'employee_id': employee_id}
            enrollment_json = json.dumps(enrollment)
            connection.request('POST', '/enrollments', enrollment_json, headers)
            response: str = connection.getresponse()
            if response.status == 200:
                created_enrollment_id: str = response.read().decode('utf-8').replace('[', '').replace(']', '').replace('"', '').strip()
                employee_id_url = f"/enrollments/{employee_id}".strip()
                connection.request('GET', employee_id_url)
                response = connection.getresponse()
                enrollments_ids: list[str] = response.read().decode('utf-8').replace('[', '').replace(']', '').replace('"', '').strip().split(',')[0::5]
                for enrollment_id in enrollments_ids:
                    if created_enrollment_id == enrollment_id:
                        created_enrollment_id_url = f"/enrollments/{created_enrollment_id}".strip()
                        # print(created_enrollment_id_url)
                        enrollment_status: dict[str, str] = {'status': 'In progress'}
                        enrollment_status_json = json.dumps(enrollment_status)
                        # print(enrollment_status_json)
                        connection.request('PUT', created_enrollment_id_url, enrollment_status_json, headers)
                        connection.getresponse()
                        # print('PUT request successful')
                    else:
                        continue
                        # print('Enrollments dont match')
            else:
                continue


def main() -> None:
    customer_emulator()


if __name__ == '__main__':
    main()
