import random

def customer_emulator() -> None:
    customer_id_index: int = random.randint(0, len(customers_ids))
    course_id_index: int = random.randint(0, len(courses_ids))