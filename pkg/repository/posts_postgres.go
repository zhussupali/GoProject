package repository

import (
	"fmt"
	"twittie"

	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (r *PostsPostgres) Create(userId int, post twittie.Post) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (text, user_id) VALUES ($1, $2) RETURNING id", postsTable)
	row := tx.QueryRow(createListQuery, post.Text, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostsPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2",
		postsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}