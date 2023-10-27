package domain

type User struct {
	ID   int64  `json:"id" redis:"id"`
	Name string `json:"name" redis:"name"`
	Age  uint8  `json:"age" redis:"age"`
}
