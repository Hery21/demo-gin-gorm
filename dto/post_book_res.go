package dto

import " hery-ciaputra/demo-gin/models"

type BookRes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Cover       string `json:"cover"`
	AuthorID    int    `json:"authorID"`
}

func (_ *BookRes) FromBook(b *models.Book) *BookRes {
	return &BookRes{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Quantity:    b.Quantity,
		Cover:       b.Cover,
		AuthorID:    b.AuthorID,
	}
}
