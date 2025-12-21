import csv


def read_ids() -> list[int]:
    with open('/var/tmp/answers.csv', 'rt', encoding='utf-8', newline='') as answers_csv:
        reader: _csv.reader = csv.reader(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        answers_ids: list[int] = [row[0] for row in reader]
        return answers_ids


def read_created_ats() -> list[str]:
    with open('/var/tmp/answers.csv', 'rt', encoding='utf-8', newline='') as answers_csv:
        reader: _csv.reader = csv.reader(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        created_ats: list[str] = [row[2] for row in reader]
        return created_ats


def read_edited_ats() -> list[str]:
    with open('/var/tmp/answers.csv', 'rt', encoding='utf-8', newline='') as answers_csv:
        reader: _csv.reader = csv.reader(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        edited_ats: list[int] = [row[3] for row in reader]
        return edited_ats


def read_is_corrects() -> list[bool]:
    with open('/var/tmp/answers.csv', 'rt', encoding='utf-8', newline='') as answers_csv:
        reader: _csv.reader = csv.reader(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        is_corrects: list[bool] = [row[4] for row in reader]
        return is_corrects


def read_questions_ids() -> list[str]:
    with open('/var/tmp/answers.csv', 'rt', encoding='utf-8', newline='') as answers_csv:
        reader: _csv.reader = csv.reader(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        quizes_ids: list[int] = [row[5] for row in reader]
        return quizes_ids


def create_simplified_answers_csv() -> None:
    with open('/var/tmp/answers-simplified.csv', 'wt', encoding='utf-8', newline='') as answers_simplified_csv:
        answers_simplified_csv.write('id,created_at,edited_at,is_correct,question_id' + '\n')
        writer: _csv.writer = csv.writer(answers_simplified_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(read_ids(), read_created_ats(), read_edited_ats(), read_is_corrects(), read_questions_ids()):
            writer.writerow(record)


def main() -> None:
    create_simplified_answers_csv()


if __name__ == '__main__':
    main()
