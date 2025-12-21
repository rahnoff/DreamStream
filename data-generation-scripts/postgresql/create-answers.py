import csv
import datetime
import random

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def read_questions_ids() -> list[str]:
    with open('/var/tmp/questions.csv', 'rt', newline='') as questions_csv:
        reader: _csv.reader = csv.reader(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        questions_ids: list[str] = [row[0] for row in reader]
        return questions_ids


def create_answers_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    questions_ids: list[str] = read_questions_ids()
    ids_1: list[int] = [id for id in range(1, len(questions_ids) + 1)]
    ids_2: list[int] = [id for id in range(len(questions_ids) + 1, (len(questions_ids) * 2) + 1)]
    ids_3: list[int] = [id for id in range((len(questions_ids) * 2) + 1, (len(questions_ids) * 3) + 1)]
    ids_4: list[int] = [id for id in range((len(questions_ids) * 3) + 1, (len(questions_ids) * 4) + 1)]
    contents_1 = [fake.text(max_nb_chars=50) for content in range(len(ids_1))]
    contents_2 = [fake.text(max_nb_chars=50) for content in range(len(ids_2))]
    contents_3 = [fake.text(max_nb_chars=50) for content in range(len(ids_3))]
    contents_4 = [fake.text(max_nb_chars=50) for content in range(len(ids_4))]
    created_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_1]
    created_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_2]
    created_ats_3: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_3]
    created_ats_4: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_4]
    edited_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_1]
    edited_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_2]
    edited_ats_3: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_3]
    edited_ats_4: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_4]
    is_corrects_1: list[bool] = [fake.pybool(truth_probability=50) for is_correct in range(len(ids_1))]
    is_corrects_2: list[bool] = [fake.pybool(truth_probability=50) for is_correct in range(len(ids_2))]
    is_corrects_3: list[bool] = [fake.pybool(truth_probability=50) for is_correct in range(len(ids_3))]
    is_corrects_4: list[bool] = [fake.pybool(truth_probability=50) for is_correct in range(len(ids_4))]
    with open('/var/tmp/answers.csv', 'wt', encoding='utf-8', newline='') as answers_csv:
        answers_csv.write('id,content,created_at,edited_at,is_correct,question_id' + '\n')
        writer: _csv.writer = csv.writer(answers_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids_1, contents_1, created_ats_1, edited_ats_1, is_corrects_1, questions_ids):
            writer.writerow(record)
        for record in zip(ids_2, contents_2, created_ats_2, edited_ats_2, is_corrects_2, questions_ids):
            writer.writerow(record)
        for record in zip(ids_3, contents_3, created_ats_3, edited_ats_3, is_corrects_3, questions_ids):
            writer.writerow(record)
        for record in zip(ids_4, contents_4, created_ats_4, edited_ats_4, is_corrects_4, questions_ids):
            writer.writerow(record)


def main() -> None:
    create_answers_csv()


if __name__ == '__main__':
    main()
