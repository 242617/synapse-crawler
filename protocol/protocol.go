package protocol

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidAnswer = errors.New("invalid answer")
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data,omitempty"`
}
