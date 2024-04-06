package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

type MovieModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
		INSERT INTO movies (isbn, title)
		VALUES ($1, $2)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, movie.Isbn, movie.Title).Scan(&movie.ID)
}

func (m MovieModel) Get(id string) (*Movie, error) {
	query := `
		SELECT id, isbn, title
		FROM movies
		WHERE id = $1
	`

	var movie Movie

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&movie.ID, &movie.Isbn, &movie.Title)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (m MovieModel) Update(movie *Movie) error {
	query := `
		UPDATE movies
		SET isbn = $1, title = $2
		WHERE id = $3
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, movie.Isbn, movie.Title, movie.ID)
	return err
}

func (m MovieModel) Delete(id string) error {
	query := `
		DELETE FROM movies
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
