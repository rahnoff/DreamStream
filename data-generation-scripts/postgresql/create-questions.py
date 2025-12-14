import csv
import datetime
import random

import faker


def create_faker_instance() -> faker.proxy.Faker:
    fake: faker.proxy.Faker = faker.Faker()
    return fake


def read_quizes_ids() -> list[str]:
    with open('/var/tmp/quizes.csv', 'rt', newline='') as quizes_csv:
        reader: _csv.reader = csv.reader(quizes_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        next(reader)
        quizes_ids: list[str] = [row[0] for row in reader]
        return quizes_ids


def create_questions_csv() -> None:
    fake: faker.proxy.Faker = create_faker_instance()
    quizes_ids: list[str] = read_quizes_ids()
    ids_1: list[int] = [id for id in range(1, len(quizes_ids) + 1)]
    ids_2: list[int] = [id for id in range(len(quizes_ids) + 1, (len(quizes_ids) * 2) + 1)]
    ids_3: list[int] = [id for id in range((len(quizes_ids) * 2) + 1, (len(quizes_ids) * 3) + 1)]
    ids_4: list[int] = [id for id in range((len(quizes_ids) * 3) + 1, (len(quizes_ids) * 4) + 1)]
    ids_5: list[int] = [id for id in range((len(quizes_ids) * 4) + 1, (len(quizes_ids) * 5) + 1)]
    ids_6: list[int] = [id for id in range((len(quizes_ids) * 5) + 1, (len(quizes_ids) * 6) + 1)]
    ids_7: list[int] = [id for id in range((len(quizes_ids) * 6) + 1, (len(quizes_ids) * 7) + 1)]
    ids_8: list[int] = [id for id in range((len(quizes_ids) * 7) + 1, (len(quizes_ids) * 8) + 1)]
    ids_9: list[int] = [id for id in range((len(quizes_ids) * 8) + 1, (len(quizes_ids) * 9) + 1)]
    ids_10: list[int] = [id for id in range((len(quizes_ids) * 9) + 1, (len(quizes_ids) * 10) + 1)]
    contents_1 = [fake.text(max_nb_chars=50) for content in range(len(ids_1))]
    contents_2 = [fake.text(max_nb_chars=50) for content in range(len(ids_2))]
    contents_3 = [fake.text(max_nb_chars=50) for content in range(len(ids_3))]
    contents_4 = [fake.text(max_nb_chars=50) for content in range(len(ids_4))]
    contents_5 = [fake.text(max_nb_chars=50) for content in range(len(ids_5))]
    contents_6 = [fake.text(max_nb_chars=50) for content in range(len(ids_6))]
    contents_7 = [fake.text(max_nb_chars=50) for content in range(len(ids_7))]
    contents_8 = [fake.text(max_nb_chars=50) for content in range(len(ids_8))]
    contents_9 = [fake.text(max_nb_chars=50) for content in range(len(ids_9))]
    contents_10 = [fake.text(max_nb_chars=50) for content in range(len(ids_10))]
    created_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_1]
    created_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_2]
    created_ats_3: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_3]
    created_ats_4: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_4]
    created_ats_5: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_5]
    created_ats_6: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_6]
    created_ats_7: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_7]
    created_ats_8: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_8]
    created_ats_9: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_9]
    created_ats_10: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_10]
    edited_ats_1: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_1]
    edited_ats_2: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_2]
    edited_ats_3: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_3]
    edited_ats_4: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_4]
    edited_ats_5: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_5]
    edited_ats_6: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_6]
    edited_ats_7: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_7]
    edited_ats_8: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_8]
    edited_ats_9: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_9]
    edited_ats_10: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for id in ids_10]
    with open('/var/tmp/questions.csv', 'wt', encoding='utf-8', newline='') as questions_csv:
        questions_csv.write('id,content,created_at,edited_at,quiz_id' + '\n')
        writer: _csv.writer = csv.writer(questions_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids_1, contents_1, created_ats_1, edited_ats_1, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_2, contents_2, created_ats_2, edited_ats_2, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_3, contents_3, created_ats_3, edited_ats_3, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_4, contents_4, created_ats_4, edited_ats_4, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_5, contents_5, created_ats_5, edited_ats_5, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_6, contents_6, created_ats_6, edited_ats_6, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_7, contents_7, created_ats_7, edited_ats_7, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_8, contents_8, created_ats_8, edited_ats_8, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_9, contents_9, created_ats_9, edited_ats_9, quizes_ids):
            writer.writerow(record)
        for record in zip(ids_10, contents_10, created_ats_10, edited_ats_10, quizes_ids):
            writer.writerow(record)


def main() -> None:
    create_questions_csv()


if __name__ == '__main__':
    main()
