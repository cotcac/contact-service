package models

// Contact type
type Contact struct {
	ID      int    `json:"id"`
	Email   string `json:"email" binding:"required,email"`
	Name    string `json:"name" binding:"required,min=3,max=155"`
	Content string `json:"content" binding:"required,min=3"`
	Date    string `json:"date"`
}
