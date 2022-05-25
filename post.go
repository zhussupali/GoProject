package twittie

type Post struct {
	Id int `json:"id"`
	Text string `json:"text" binding:"required"`
	UserId int `json:"userId"`
}

