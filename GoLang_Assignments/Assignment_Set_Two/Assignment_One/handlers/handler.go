
package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "time"
    "Assignment_One/models"
)

type Handler struct {
    db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
    return &Handler{db: db}
}

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var blog models.Blog
    if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    blog.Timestamp = time.Now()
    
    result, err := h.db.Exec(
        "INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)",
        blog.Title, blog.Content, blog.Author, blog.Timestamp,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    blog.ID = id

    json.NewEncoder(w).Encode(blog)
}

func (h *Handler) HandleBlog(w http.ResponseWriter, r *http.Request) {
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) != 3 {
        http.Error(w, "Invalid path", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(pathParts[2], 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        h.getBlog(w, r, id)
    case http.MethodPut:
        h.updateBlog(w, r, id)
    case http.MethodDelete:
        h.deleteBlog(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *Handler) getBlog(w http.ResponseWriter, r *http.Request, id int64) {
    var blog models.Blog
    err := h.db.QueryRow(
        "SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?",
        id,
    ).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)

    if err == sql.ErrNoRows {
        http.Error(w, "Blog not found", http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(blog)
}

func (h *Handler) updateBlog(w http.ResponseWriter, r *http.Request, id int64) {
    var blog models.Blog
    if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := h.db.Exec(
        "UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?",
        blog.Title, blog.Content, blog.Author, id,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    blog.ID = id
    json.NewEncoder(w).Encode(blog)
}

func (h *Handler) deleteBlog(w http.ResponseWriter, r *http.Request, id int64) {
    result, err := h.db.Exec("DELETE FROM blogs WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Blog not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListBlogs(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    rows, err := h.db.Query("SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var blogs []models.Blog
    for rows.Next() {
        var blog models.Blog
        err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        blogs = append(blogs, blog)
    }

    json.NewEncoder(w).Encode(blogs)
}