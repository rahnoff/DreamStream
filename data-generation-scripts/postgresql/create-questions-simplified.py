import csv


def read_ids() -> list[int]:
    with open('/var/tmp/questions.csv', 'rt', encoding='utf-8', newline='') as questions_csv:
        reader: _csv.reader = csv.reader(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        questions_ids: list[int] = [row[0] for row in reader]
        return questions_ids


def read_created_ats() -> list[str]:
    with open('/var/tmp/questions.csv', 'rt', encoding='utf-8', newline='') as questions_csv:
        reader: _csv.reader = csv.reader(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        created_ats: list[str] = [row[2] for row in reader]
        return created_ats


def read_edited_ats() -> list[str]:
    with open('/var/tmp/questions.csv', 'rt', encoding='utf-8', newline='') as questions_csv:
        reader: _csv.reader = csv.reader(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        edited_ats: list[int] = [row[3] for row in reader]
        return edited_ats


def read_quizes_ids() -> list[str]:
    with open('/var/tmp/questions.csv', 'rt', encoding='utf-8', newline='') as questions_csv:
        reader: _csv.reader = csv.reader(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        quizes_ids: list[int] = [row[4] for row in reader]
        return quizes_ids


def create_simplified_questions_csv() -> None:
    with open('/var/tmp/questions-simplified.csv', 'wt', encoding='utf-8', newline='') as questions_simplified_csv:
        questions_simplified_csv.write('id,created_at,edited_at,quiz_id' + '\n')
        writer: _csv.writer = csv.writer(questions_simplified_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(read_ids(), read_created_ats(), read_edited_ats(), read_quizes_ids()):
            writer.writerow(record)


def main() -> None:
    create_simplified_questions_csv()


if __name__ == '__main__':
    main()
