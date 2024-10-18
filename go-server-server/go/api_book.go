package swagger

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	bookStore = make(map[string]Book)
	mutex     sync.Mutex
)

// CreateBook adds a new book to the bookStore
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Check if the book ID already exists
	if _, exists := bookStore[book.BookID]; exists {
		http.Error(w, "Book with the same ID already exists", http.StatusConflict)
		return
	}

	// Add the book to the store
	bookStore[book.BookID] = book
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook removes a book by ID from the bookStore
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookID"]

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := bookStore[bookID]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Delete the book
	delete(bookStore, bookID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

// GetBookById retrieves a specific book by ID from the bookStore
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookID"]

	mutex.Lock()
	defer mutex.Unlock()

	book, exists := bookStore[bookID]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return the book
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// ListBooks returns all the books in the bookStore
func ListBooks(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	books := make([]Book, 0, len(bookStore))
	for _, book := range bookStore {
		books = append(books, book)
	}

	// Return the list of books
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// UpdateBookById updates a book's details by ID
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookID"]

	var updatedBook Book

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Check if the book exists
	if _, exists := bookStore[bookID]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update the book's details
	updatedBook.BookID = bookID
	bookStore[bookID] = updatedBook

	// Return the updated book
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}
