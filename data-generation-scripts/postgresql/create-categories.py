import csv
import datetime


def create_categories_csv() -> None:
    names: list[str] = ['Finance', 'HR', 'IT', 'Management', 'Marketing', 'Statistics']
    ids: list[int] = [id for id in list(range(1, len(names) + 1))]
    created_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for created_at in range(len(names))]
    edited_ats: list[str] = [datetime.datetime.now().astimezone().isoformat().replace('T', ' ') for edited_at in range(len(names))]
    with open('/var/tmp/categories.csv', 'wt', encoding='utf-8', newline='') as categories_csv:
        categories_csv.write('id,created_at,edited_at,name' + '\n')
        writer: _csv.writer = csv.writer(categories_csv, dialect='unix', delimiter=',', escapechar='\\', quoting=csv.QUOTE_NONE)
        for record in zip(ids, created_ats, edited_ats, names):
            writer.writerow(record)


def main() -> None:
    create_categories_csv()


if __name__ == '__main__':
    main()
