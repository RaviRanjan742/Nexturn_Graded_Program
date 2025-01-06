from app import create_app, db
from app.models import Book

app = create_app()

with app.app_context():
    db.create_all()

    sample_books = [
        {"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "published_year": 1925, "genre": "Fiction"},
        {"title": "To Kill a Mockingbird", "author": "Harper Lee", "published_year": 1960, "genre": "Fiction"}
    ]

    for book_data in sample_books:
        book = Book(**book_data)
        db.session.add(book)
    db.session.commit()
