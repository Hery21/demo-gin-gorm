package dto

type BookReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	Cover       string `json:"cover"`
	AuthorID    int    `json:"authorID" binding:"required"`
}
