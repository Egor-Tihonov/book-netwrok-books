package model

type Book struct {
	Author
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
