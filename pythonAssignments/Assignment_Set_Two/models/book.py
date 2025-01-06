from dataclasses import dataclass
from typing import Dict, List

@dataclass
class Book:
    title: str
    author: str
    price: float
    quantity: int

    def display(self) -> str:
        return f"Title: {self.title}\nAuthor: {self.author}\nPrice: ${self.price:.2f}\nQuantity: {self.quantity}"

    @staticmethod
    def validate_price(price: float) -> bool:
        return price > 0

    @staticmethod
    def validate_quantity(quantity: int) -> bool:
        return quantity >= 0