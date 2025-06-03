import csv

import faker


faker: faker.proxy.Faker = faker.Faker()
first_names: list[str] = [faker.first_name() for first_name in range(10000)]
last_names: list[str] = [faker.last_name() for last_name in range(10000)]
ids: list[str] = [faker.uuid4() for id in range(10000)]
