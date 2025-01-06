from dataclasses import dataclass
from datetime import datetime
from .customer import Customer

@dataclass
class Transaction(Customer):
    book_title: str
    quantity_sold: int
    transaction_date: datetime = None

    def __post_init__(self):
        if self.transaction_date is None:
            self.transaction_date = datetime.now()

    def display(self) -> str:
        return f"{super().display()}\nBook: {self.book_title}\nQuantity: {self.quantity_sold}\nDate: {self.transaction_date}"