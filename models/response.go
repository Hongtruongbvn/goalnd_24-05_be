package models

// MessageResponse dùng cho phản hồi thành công dạng {"message": "..."}
type MessageResponse struct {
	Message string `json:"message" example:"Thành công"`
}

// ErrorResponse dùng cho phản hồi lỗi dạng {"error": "..."}
type ErrorResponse struct {
	Error string `json:"error" example:"Có lỗi xảy ra"`
}
