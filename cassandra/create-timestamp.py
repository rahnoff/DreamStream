import datetime


def generate_example_data() -> str:
    timestamp: str = datetime.datetime.now().astimezone().isoformat().replace('T', ' ')
    return timestamp
