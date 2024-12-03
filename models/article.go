package models

type Article struct {
	ArticleID  string `json:"article_id"`
	CategoryID string `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   string `json:"author_id"`
	Slug       string `json:"slug"`
}
