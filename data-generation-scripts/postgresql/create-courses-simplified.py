import csv


def read_ids() -> list[int]:
    with open('/var/tmp/courses.csv', 'rt', encoding='utf-8', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        courses_ids: list[int] = [row[0] for row in reader]
        return courses_ids


def read_created_ats() -> list[str]:
    with open('/var/tmp/courses.csv', 'rt', encoding='utf-8', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        created_ats: list[str] = [row[2] for row in reader]
        return created_ats


def read_edited_ats() -> list[str]:
    with open('/var/tmp/courses.csv', 'rt', encoding='utf-8', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        edited_ats: list[int] = [row[4] for row in reader]
        return edited_ats


def read_names() -> list[str]:
    with open('/var/tmp/courses.csv', 'rt', encoding='utf-8', newline='') as courses_csv:
        reader: _csv.reader = csv.reader(courses_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        names: list[int] = [row[8] for row in reader]
        return names


def create_simplified_courses_csv() -> None:
    with open('/var/tmp/courses-simplified.csv', 'wt', encoding='utf-8', newline='') as courses_simplified_csv:
        courses_simplified_csv.write('id,created_at,edited_at,name' + '\n')
        writer: _csv.writer = csv.writer(courses_simplified_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(read_ids(), read_created_ats(), read_edited_ats(), read_names()):
            writer.writerow(record)


def main() -> None:
    create_simplified_courses_csv()


if __name__ == '__main__':
    main()
