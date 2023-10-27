package dto

import "github.com/vildan-valeev/melvad_test/internal/domain"

type SignDtoRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type SignDtoResponse struct {
	Code string `json:"Code"`
}

func ToDTO(sign domain.Sign) SignDtoResponse {
	return SignDtoResponse{
		Code: sign.Hash,
	}
}
