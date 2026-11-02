package requests

type BlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Slug    string `json:"slug" binding:"required"`
}
