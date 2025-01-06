from typing import List
from models.customer import Customer

class CustomerManager:
    def __init__(self):
        self.customers: List[Customer] = []

    def add_customer(self, name: str, email: str, phone: str) -> bool:
        try:
            if not Customer.validate_email(email):
                raise ValueError("Invalid email format")
            if not Customer.validate_phone(phone):
                raise ValueError("Invalid phone number (must be 10 digits)")
            
            customer = Customer(name, email, phone)
            self.customers.append(customer)
            return True
        except ValueError as e:
            print(f"Error: {e}")
            return False

    def view_customers(self) -> List[Customer]:
        return self.customers