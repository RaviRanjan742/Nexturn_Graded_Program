VALID_GENRES = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi", "Fantasy"]

def validate_book_data(data):
    if not all(key in data for key in ["title", "author", "published_year", "genre"]):
        return "Missing required fields"
    if not isinstance(data['published_year'], int) or data['published_year'] <= 0:
        return "Invalid published_year"
    if data['genre'] not in VALID_GENRES:
        return f"Invalid genre. Valid genres are: {', '.join(VALID_GENRES)}"
    return None
