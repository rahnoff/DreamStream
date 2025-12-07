import csv
import datetime
import random

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def read_courses_ids() -> list[str]:
    with open('/var/tmp/courses.csv', 'rt', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        courses_ids: list[str] = [row[0] for row in reader]
        return courses_ids


def create_quizes_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    courses_ids: list[str] = read_courses_ids()
    ids: list[int] = [id for id in range(1, len(courses_ids) * 2)]
    courses_ids_randomized: list[str] = [random.choice(courses_ids) for course_id in courses_ids]
    ids_1: list[str] = [fake.uuid4() for course_id in courses_ids]
    ids_2: list[str] = [fake.uuid4() for course_id in courses_ids]
    created_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for course_id in courses_ids]
    created_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for course_id in courses_ids]
    edited_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for course_id in courses_ids]
    edited_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for course_id in courses_ids]
    with open('quizes.csv', 'wt', encoding='utf-8', newline='') as quizes_csv:
        quizes_csv.write('id,course_id,created_at,edited_at' + '\n')
        writer: _csv.writer = csv.writer(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids_1, courses_ids, created_ats_1, edited_ats_1):
            writer.writerow(record)
        for record in zip(ids_2, courses_ids, created_ats_2, edited_ats_2):
            writer.writerow(record)


def main() -> None:
    create_quizes_csv()


if __name__ == '__main__':
    main()