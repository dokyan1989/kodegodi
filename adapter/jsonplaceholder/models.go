package jsonplaceholder

type Post struct {
	UserID uint64 `json:"userId"`
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
