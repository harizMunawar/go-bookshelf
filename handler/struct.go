package handler

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type shortDetail struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	BookId int    `json:"bookId"`
}
