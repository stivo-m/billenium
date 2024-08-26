package dto

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type LoginDto struct {
	Token string `validate:"required,min=10" json:"token"`
	Type  string `validate:"required,oneof=google apple" json:"type"`
}

type InsertTxnDto struct {
	Type   string `json:"type" validate:"required,oneof=incoming outgoing"`
	Name   string `json:"name" validate:"required,min=4"`
	Amount string `json:"amount" validate:"required,numeric,min=1"`
}
