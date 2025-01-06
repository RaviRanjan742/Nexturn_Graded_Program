from dataclasses import dataclass
import re
from typing import Dict, List

@dataclass
class Customer:
    name: str
    email: str
    phone: str

    def display(self) -> str:
        return f"Name: {self.name}\nEmail: {self.email}\nPhone: {self.phone}"

    @staticmethod
    def validate_email(email: str) -> bool:
        pattern = r'^[\w\.-]+@[\w\.-]+\.\w+$'
        return bool(re.match(pattern, email))

    @staticmethod
    def validate_phone(phone: str) -> bool:
        pattern = r'^\d{10}$'
        return bool(re.match(pattern, phone))