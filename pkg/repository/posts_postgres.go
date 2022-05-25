package repository

import (
	"fmt"
	"strings"
	"twittie"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *PostsPostgres) GetAll(userId int) ([]twittie.Post, error) {
	var posts []twittie.Post

	query := fmt.Sprintf("SELECT id, text, user_id from %s WHERE user_id = $1",
		postsTable)
	err := r.db.Select(&posts, query, userId)

	return posts, err
}

func (r *PostsPostgres) GetById(userId, postId int) (twittie.Post, error) {
	var list twittie.Post

	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1 AND id = $2`,
		postsTable)
	err := r.db.Get(&list, query, userId, postId)

	return list, err
}


func (r *PostsPostgres) Delete(userId, postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2",
		postsTable)
	_, err := r.db.Exec(query, userId, postId)

	return err
}


func (r *PostsPostgres) Update(userId, postId int, input twittie.UpdatePostInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}
	
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d",
		postsTable, setQuery, argId, argId+1)
	args = append(args, postId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}