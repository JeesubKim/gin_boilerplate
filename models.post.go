package main

type post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getAllPost() []post {
	return []post{{0, "title1", "content1"}, {1, "title2", "content2"}}
}
