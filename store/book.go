package store

type Book struct {
	ID 		 int 		`json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (s *Store) GetAllBooks() ([]Book, error) {
	rows, err := s.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (s *Store) CreateBook(b *Book) error {
	_, err := s.db.Exec("INSERT INTO books (title, author) VALUES ($1, $2)", b.Title, b.Author)
	return err
}

func (s *Store) UpdateBook(b *Book) error {
	_, err := s.db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3", b.Title, b.Author, b.ID)
	return err
}

func (s *Store) DeleteBook(id int) error {
	_, err := s.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
