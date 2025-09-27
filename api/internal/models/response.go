package models

type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Message string      `json:"message,omitempty"`
}

// ErrorResponse resposta específica para erros
type ErrorResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
    Code    int    `json:"code"`
}

// SuccessResponse resposta específica para sucesso
type SuccessResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Message string      `json:"message,omitempty"`
}
