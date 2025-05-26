import csv
import datetime
import uuid


NUMBER_OF_ROWS: int = 1000000
AUTHORS: tuple[str] = ('Steven Baker', 'Denise Carr', 'Michelle Robinson', 'Kristen Terry', 'Christopher Walsh')

def create_timestamps(quantity: int) -> datetime.datetime:
    # timestamp_without_timezone: datetime.datetime = datetime.datetime.now()
    # timestamp_with_timezone: datetime.datetime = timestamp_without_timezone.astimezone()
    # timestamps_with_timezones = [timestamp_without_timezone.astimezone() for iteration in range(quantity)]
    # timestamp: str = datetime.datetime.now().astimezone().isoformat().replace('T', ' ')
    # print(timestamp_with_timezone)
    # return timestamps_with_timezones
    # print(timestamps_with_timezones)
    # timestamps: list[datetime] = [(fake.date_time_between_dates(datetime_start=datetime.combine(date(date.today().year, start_month, 1),time(hour=0, minute=0, second=0)),datetime_end=datetime.combine(date(date.today().year, end_month, 31),time(hour=23, minute=59, second=59),),)).isoformat(sep=" ")for timestamp in range(quantity)]
    # return timestamps
    


def create_uuid() -> uuid.UUID:
    uuid4: uuid.UUID = uuid.uuid4()
    return uuid4


def create_csv() -> None:
    for iteration in range(100):
        timestamp_with_timezone = create_timestamp()
        uuid4 = create_uuid

def write_csv(file_name: str, *args) -> None:
    with open(file_name, "w", newline="") as csv_file:
        writer = csv.writer(csv_file, dialect="unix", delimiter=",", escapechar="\\", quoting=csv.QUOTE_NONE)
        for record in zip():
            writer.writerow(record)


def main() -> None:
    create_timestamps(10)


if __name__ == '__main__':
    main()