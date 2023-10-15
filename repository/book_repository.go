package repository

import (
	"context"
	"errors"

	"github.com/brandoniu/lets-go/models"
	"github.com/jackc/pgx/v4"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	FindByID(id int64) (*models.Book, error)
	Create(b *models.Book) (int64, error)
	Update(b *models.Book) error
	Delete(id int64) error
}

type PostgresBookRepostiory struct {
	conn *pgx.Conn
}

func NewPostgresBookRepository(connString string) (*PostgresBookRepostiory, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &PostgresBookRepostiory{conn: conn}, nil
}

func (p *PostgresBookRepostiory) FindAll() ([]models.Book, error) {
	books := []models.Book{}
	query := `SELECT id, title, author FROM books`
	rows, err := p.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (p *PostgresBookRepostiory) FindByID(id int64) (*models.Book, error) {
	query := `SELECT id, title, author FROM books WHERE id=$1`
	row := p.conn.QueryRow(context.Background(), query, id)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Author); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (p *PostgresBookRepostiory) Create(b *models.Book) (int64, error) {
	query := `INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id`
	var id int64
	if err := p.conn.QueryRow(context.Background(), query, b.Title, b.Author).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresBookRepostiory) Update(b *models.Book) error {
	query := `UPDATE books SET title=$1, author=$2, published_date=$3 WHERE id=$4`
	_, err := p.conn.Exec(context.Background(), query, b.Title, b.Author, b.ID)
	return err
}

func (p *PostgresBookRepostiory) Delete(id int64) error {
	query := `DELETE FROM books WHERE id=$1`
	_, err := p.conn.Exec(context.Background(), query, id)
	return err
}
