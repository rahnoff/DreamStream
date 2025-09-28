import csv
import datetime
import random

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def read_categories_ids() -> list[str]:
    with open('categories.csv', 'rt', newline='') as csv_file:
        reader: _csv.reader = csv.reader(csv_file, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        categories_ids: list[str] = [row[0] for row in reader]
        return categories_ids


def create_courses_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    ids: list[str] = [fake.uuid4() for id in range(1000)]
    categories_ids: list[str] = [random.choice(read_categories_ids()) for category_id in range(1000)]
    created_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for created_at in range(1000)]
    edited_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for edited_at in range(1000)]
    names: list[str] = [fake.text(max_nb_chars=20) for name in range(1000)]
    with open('courses.csv', 'wt', encoding='utf-8', newline='') as courses_csv:
        courses_csv.write('id,category_id,created_at,edited_at,name' + '\n')
        writer: _csv.writer = csv.writer(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids, categories_ids, created_ats, edited_ats, names):
            writer.writerow(record)


def main() -> None:
    create_courses_csv()


if __name__ == '__main__':
    main()
