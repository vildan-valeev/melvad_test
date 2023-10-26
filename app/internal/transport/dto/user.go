package dto

type UserUpdateDtoRequest struct {
	ID    int64  `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value uint8  `json:"value,omitempty"`
}

type UserUpdateDtoResponse struct {
	Value uint8 `json:"value"`
}

type UserCreateDtoRequest struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type UserCreateDtoResponse struct {
	ID int64 `json:"id"`
}
