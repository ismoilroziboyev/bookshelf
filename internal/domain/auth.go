package domain

type SignUpPayload struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Key    string `json:"key" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Key    string `json:"key"`
	Email  string `json:"email"`
}
