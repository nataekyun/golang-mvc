package message

type UserRequest struct {
	UserID string `json:"userID"`
}

type UserResponse struct {
	UserID    int64  `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
