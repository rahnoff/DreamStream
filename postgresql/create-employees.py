import csv

import faker


faker: faker.proxy.Faker = faker.Faker()
first_names: list[str] = [faker.first_name() for first_name in range(10000)]
last_names: list[str] = [faker.last_name() for last_name in range(10000)]
ids: list[str] = [faker.uuid4() for id in range(10000)]
with open('employees.csv', 'w', newline='') as employees_csv:
    writer: _csv.writer = csv.writer(employees_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
    for record in zip(ids, first_names, last_names):
        writer.writerow(record)
