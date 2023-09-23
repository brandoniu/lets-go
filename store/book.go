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