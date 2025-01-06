from flask import Blueprint, request, jsonify
from . import db
from .models import Book
from .utils import validate_book_data

book_blueprint = Blueprint('books', __name__)

@book_blueprint.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    error = validate_book_data(data)
    if error:
        return jsonify({"error": "Invalid data", "message": error}), 400

    new_book = Book(**data)
    db.session.add(new_book)
    db.session.commit()
    return jsonify({"message": "Book added successfully", "book_id": new_book.id}), 201

@book_blueprint.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    return jsonify([book.to_dict() for book in books]), 200

@book_blueprint.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    return jsonify(book.to_dict()), 200

@book_blueprint.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    data = request.get_json()
    error = validate_book_data(data)
    if error:
        return jsonify({"error": "Invalid data", "message": error}), 400

    book.title = data['title']
    book.author = data['author']
    book.published_year = data['published_year']
    book.genre = data['genre']
    db.session.commit()
    return jsonify({"message": "Book updated successfully"}), 200

@book_blueprint.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    db.session.delete(book)
    db.session.commit()
    return jsonify({"message": "Book deleted successfully"}), 200
