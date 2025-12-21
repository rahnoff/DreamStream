import csv
import datetime

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def create_employees_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    ids: list[str] = [fake.uuid4() for id in range(3000000)]
    created_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for created_at in range(3000000)]
    edited_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for edited_at in range(3000000)]
    emails: list[str] = [fake.unique.email(safe=True, domain='contoso.com') for email in range(3000000)]
    first_names: list[str] = [fake.first_name() for first_name in range(3000000)]
    last_names: list[str] = [fake.last_name() for last_name in range(3000000)]
    with open('/var/tmp/employees.csv', 'wt', encoding='utf-8', newline='') as employees_csv:
        employees_csv.write('id,created_at,edited_at,email,first_name,last_name' + '\n')
        writer: _csv.writer = csv.writer(employees_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids, created_ats, edited_ats, emails, first_names, last_names):
            writer.writerow(record)


def main() -> None:
    create_employees_csv()


if __name__ == '__main__':
    main()
