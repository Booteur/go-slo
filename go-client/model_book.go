package swagger

type Book struct {
	BookID string `json:"bookID,omitempty"`
	BookDescription string `json:"bookDescription,omitempty"`
	BookTitle string `json:"bookTitle,omitempty"`
	BookAuthor string `json:"bookAuthor,omitempty"`
}
