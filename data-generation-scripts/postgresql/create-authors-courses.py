import csv
import datetime
import random

import faker


def read_authors_ids() -> list[str]:
    with open('/var/tmp/authors.csv', 'rt', newline='') as authors_csv:
        reader: _csv.reader = csv.reader(authors_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        authors_ids: list[str] = [row[0] for row in reader]
        return authors_ids


def read_courses_ids() -> list[str]:
    with open('/var/tmp/courses.csv', 'rt', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        courses_ids: list[str] = [row[0] for row in reader]
        return courses_ids


def create_authors_courses_csv() -> None:
    authors_ids: list[str] = read_authors_ids()
    courses_ids: list[str] = read_courses_ids()
    authors_of_courses_ids: list[str] = [random.choice(authors_ids) for course_id in courses_ids]
    with open('/var/tmp/authors-courses.csv', 'wt', encoding='utf-8', newline='') as authors_courses_csv:
        authors_courses_csv.write('author_id,course_id' + '\n')
        writer: _csv.writer = csv.writer(authors_courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(authors_of_courses_ids, courses_ids):
            writer.writerow(record)


def main() -> None:
    create_authors_courses_csv()


if __name__ == '__main__':
    main()