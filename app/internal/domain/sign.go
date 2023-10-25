package domain

type Sign struct {
	ID   uint64 `json:"id"`
	HMAC string `json:"code"` // HMAC-SHA512 подпись.
}
