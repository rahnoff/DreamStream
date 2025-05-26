import datetime


def generate_example_data() -> None:
    timestamp: str = datetime.datetime.now().astimezone().isoformat().replace('T', ' ')
