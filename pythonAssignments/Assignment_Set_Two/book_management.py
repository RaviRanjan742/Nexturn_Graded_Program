
from typing import List, Optional
from models.book import Book

class BookManager:
    def __init__(self):
        self.books: List[Book] = []

    def add_book(self, title: str, author: str, price: float, quantity: int) -> bool:
        try:
            if not Book.validate_price(price):
                raise ValueError("Price must be positive")
            if not Book.validate_quantity(quantity):
                raise ValueError("Quantity must be non-negative")
            
            # Check if book already exists
            existing_book = self.search_book(title)
            if existing_book:
                existing_book.quantity += quantity
                return True
            
            book = Book(title, author, price, quantity)
            self.books.append(book)
            return True
        except ValueError as e:
            print(f"Error: {e}")
            return False

    def view_books(self) -> List[Book]:
        return self.books

    def search_book(self, title: str) -> Optional[Book]:
        for book in self.books:
            if book.title.lower() == title.lower():
                return book
        return None

    def update_quantity(self, title: str, quantity_sold: int) -> bool:
        book = self.search_book(title)
        if not book:
            return False
        if book.quantity < quantity_sold:
            return False
        book.quantity -= quantity_sold
        return True
