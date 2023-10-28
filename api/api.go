package api

import (
	"encoding/json"
	"net/http"
)

// CoinBalanceParams Coin Balance Struct
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse Coin Balance Response Struct
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Err Response
type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)
