import csv
import datetime
import random

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def read_categories_ids() -> list[int]:
    with open('/var/tmp/categories.csv', 'rt', encoding='utf-8', newline='') as categories_csv:
        reader: _csv.reader = csv.reader(categories_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        categories_ids: list[int] = [row[0] for row in reader]
        return categories_ids


def create_courses_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    ids: list[int] = [id for id in range(1, 100001)]
    languages: list[str] = ['EN', 'DE']
    lengths: list[int] = [20, 30, 40]
    languages_generated: list[str] = [random.choice(languages) for language in range(100000)]
    lengths_generated: list[str] = [random.choice(lengths) for length in range(100000)]
    categories_ids: list[str] = [random.choice(read_categories_ids()) for category_id in range(100000)]
    created_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for created_at in range(100000)]
    descriptions: list[str] = [fake.paragraph(nb_sentences=5) for description in range(100000)]
    edited_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for edited_at in range(100000)]
    filenames: list[str] = [fake.unique.url() for url in range(100000)]
    names: list[str] = [fake.unique.text(max_nb_chars=20) for name in range(100000)]
    with open('/var/tmp/courses.csv', 'wt', encoding='utf-8', newline='') as courses_csv:
        courses_csv.write('id,category_id,created_at,description,edited_at,filename,language,length,name' + '\n')
        writer: _csv.writer = csv.writer(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids, categories_ids, created_ats, descriptions, edited_ats, filenames, languages_generated, lengths_generated, names):
            writer.writerow(record)


def main() -> None:
    create_courses_csv()


if __name__ == '__main__':
    main()
