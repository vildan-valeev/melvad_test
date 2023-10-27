package domain

type Sign struct {
	ID   uint64 `json:"id"`
	Hash string `json:"code"` // HMAC-SHA512 подпись в Hex.
}
