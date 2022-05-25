package twittie

type Post struct {
	Id int `json:"id"`
	Text string `json:"text" binding:"required"`
	User_id int `json:"userId"`
}

