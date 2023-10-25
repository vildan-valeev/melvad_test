package domain

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}