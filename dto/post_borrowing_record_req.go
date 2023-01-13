package dto

type BorrowingRecordReq struct {
	UserID int `json:"userID" binding:"required"`
	BookID int `json:"bookID" binding:"required"`
}
