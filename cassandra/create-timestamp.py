import datetime


def create_timestamp() -> str:
    timestamp: str = datetime.datetime.now().astimezone().isoformat().replace('T', ' ')
    return timestamp
