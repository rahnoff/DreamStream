import csv
import datetime

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def create_categories_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    ids: list[str] = [fake.uuid4() for id in range(100)]
    created_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for created_at in range(100)]
    edited_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for edited_at in range(100)]
    names: list[str] = [word.capitalize() for word in fake.words(nb=100,unique=True)]
    with open('categories.csv', 'wt', encoding='utf-8', newline='') as categories_csv:
        categories_csv.write('id,created_at,edited_at,name' + '\n')
        writer: _csv.writer = csv.writer(categories_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids, created_ats, edited_ats, names):
            writer.writerow(record)


def main() -> None:
    create_categories_csv()


if __name__ == '__main__':
    main()
