package models

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GenerateError(code int, msg string) Error {
	return Error{Code: code, Message: msg}
}
