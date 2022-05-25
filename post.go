package twittie

import "errors"

type Post struct {
	Id int `json:"id"`
	Text string `json:"text" binding:"required"`
	User_id int `json:"userId"`
}

func (i UpdatePostInput) Validate() error {
	if i.Text == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdatePostInput struct {
	Text       *string `json:"text"`
}