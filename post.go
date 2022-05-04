package twittie

type Post struct {
	Id int `json:"id"`
	Text string `json:"text"`
	UserId int `json:"userId"`
}

