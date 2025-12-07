import csv
import random


def read_employees_ids() -> list[str]:
    with open('/var/tmp/employees.csv', 'rt', newline='') as employees_csv:
        reader: _csv.reader = csv.reader(employees_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        employees_ids: list[str] = [row[0] for row in reader]
        return employees_ids


def create_authors_csv() -> None:
    employees_ids: list[str] = read_employees_ids()
    ids: list[str] = [random.choice(employees_ids) for id in range(1000)]
    with open('/var/tmp/authors.csv', 'wt', encoding='utf-8', newline='') as authors_csv:
        authors_csv.write('id' + '\n')
        writer: _csv.writer = csv.writer(authors_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids):
            writer.writerow(record)


def main() -> None:
    create_authors_csv()


if __name__ == '__main__':
    main()
