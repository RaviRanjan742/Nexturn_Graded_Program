import unittest
from app import create_app, db
from app.models import Book

class BookBuddyTestCase(unittest.TestCase):
    def setUp(self):
        self.app = create_app()
        self.client = self.app.test_client()
        with self.app.app_context():
            db.create_all()

    def tearDown(self):
        with self.app.app_context():
            db.drop_all()

    def test_add_book(self):
        response = self.client.post('/books', json={
            "title": "1984",
            "author": "George Orwell",
            "published_year": 1949,
            "genre": "Fiction"
        })
        self.assertEqual(response.status_code, 201)

    # Add more tests for other endpoints

if __name__ == '__main__':
    unittest.main()
