import random


customer_id: int = random.randint(0, len(customers_ids))
course_id: int = random.randint(0, len(courses_ids))
print(customer_id, course_id)