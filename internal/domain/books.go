package domain

const (
	BookStatusNew = iota
	BookStatusReading
	BookSttausFinished
)

type CreateBookPayload struct {
	ISBN string `json:"isbn" binding:"required"`
}

type EditBookPayload struct {
	ID     int64 `json:"id"`
	Status int32 `json:"status" binding:"min=0,max=3"`
}
