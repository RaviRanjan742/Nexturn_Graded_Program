
from typing import List
from models.transaction import Transaction
from book_management import BookManager
from customer_management import CustomerManager

class SalesManager:
    def __init__(self, book_manager: BookManager, customer_manager: CustomerManager):
        self.transactions: List[Transaction] = []
        self.book_manager = book_manager
        self.customer_manager = customer_manager

    def sell_book(self, customer_name: str, customer_email: str, customer_phone: str, 
                 book_title: str, quantity: int) -> bool:
        try:
            
            book = self.book_manager.search_book(book_title)
            if not book:
                raise ValueError("Book not found")
            if book.quantity < quantity:
                raise ValueError(f"Insufficient quantity. Only {book.quantity} available")

            
            if not self.book_manager.update_quantity(book_title, quantity):
                return False

            
            transaction = Transaction(
                customer_name, customer_email, customer_phone,
                book_title, quantity
            )
            self.transactions.append(transaction)
            return True
        except ValueError as e:
            print(f"Error: {e}")
            return False

    def view_sales(self) -> List[Transaction]:
        return self.transactions