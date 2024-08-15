package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type DataTemplate struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type ErrorTemplate struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	return w
}

func SetResponse(w http.ResponseWriter, httpStatus int, data interface{}, message string, success bool) {
	w = setHeader(w)

	response := DataTemplate{
		Message: message,
		Success: success,
		Status:  http.StatusText(httpStatus),
		Data:    data,
	}

	encoded, err := json.Marshal(response)

	if err != nil {
		log.Printf("[ERROR][Util][WriteSuccess] marshalling response, %+v\n", err)
	}

	w.WriteHeader(httpStatus)
	w.Write(encoded)
}

func SetError(w http.ResponseWriter, httpStatus int, err error) {
	w = setHeader(w)

	response := ErrorMsg{
		Error: err.Error(),
	}

	encoded, err := json.Marshal(response)

	if err != nil {
		log.Printf("[ERROR][Util][WriteError] marshalling response, %+v\n", err)
	}

	w.WriteHeader(httpStatus)
	w.Write(encoded)
}

func ReturnInternalServerError(w http.ResponseWriter) {
	w = setHeader(w)
	response := ErrorMsg{
		Error: "Internal Server Error",
	}

	encoded, err := json.Marshal(response)

	if err != nil {
		log.Printf("[ERROR][Util][WriteError] marshalling response, %+v\n", err)
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(encoded)
}
