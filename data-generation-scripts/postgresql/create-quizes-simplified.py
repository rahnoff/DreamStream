import csv


def read_ids() -> list[int]:
    with open('/var/tmp/quizes.csv', 'rt', encoding='utf-8', newline='') as quizes_csv:
        reader: _csv.reader = csv.reader(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        quizes_ids: list[int] = [row[0] for row in reader]
        return quizes_ids


def read_courses_ids() -> list[str]:
    with open('/var/tmp/quizes.csv', 'rt', encoding='utf-8', newline='') as quizes_csv:
        reader: _csv.reader = csv.reader(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        courses_ids: list[str] = [row[1] for row in reader]
        return courses_ids


def read_created_ats() -> list[str]:
    with open('/var/tmp/quizes.csv', 'rt', encoding='utf-8', newline='') as quizes_csv:
        reader: _csv.reader = csv.reader(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        created_ats: list[str] = [row[2] for row in reader]
        return created_ats


def read_edited_ats() -> list[str]:
    with open('/var/tmp/quizes.csv', 'rt', encoding='utf-8', newline='') as quizes_csv:
        reader: _csv.reader = csv.reader(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        edited_ats: list[int] = [row[3] for row in reader]
        return edited_ats


def create_simplified_quizes_csv() -> None:
    with open('/var/tmp/quizes-simplified.csv', 'wt', encoding='utf-8', newline='') as quizes_simplified_csv:
        quizes_simplified_csv.write('id,course_id,created_at,edited_at' + '\n')
        writer: _csv.writer = csv.writer(quizes_simplified_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(read_ids(), read_courses_ids(), read_created_ats(), read_edited_ats()):
            writer.writerow(record)


def main() -> None:
    create_simplified_quizes_csv()


if __name__ == '__main__':
    main()
