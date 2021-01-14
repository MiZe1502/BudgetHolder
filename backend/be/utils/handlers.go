package utils

import (
	"encoding/json"
	"net/http"
)

// SliceData stores data to get slice of some entities from db
type SliceData struct {
	From  int `json:"from"`
	Count int `json:"count"`
}

type TableData struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// Message creates message for handlers
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message}
}

// MessageData creates message with data field
func MessageData(msg map[string]interface{}, data []byte) map[string]interface{} {
	msg["data"] = string(data)
	return msg
}

// MessageError wraps message as error message for handlers
func MessageError(msg map[string]interface{}, errorCode int) map[string]interface{} {
	msg["errorCode"] = errorCode
	return msg
}

func hideError(data map[string]interface{}) map[string]interface{} {
	data["message"] = ""
	data["errorCode"] = http.StatusBadRequest

	return data
}

//RespondError logs error and hides it with 400 for response
func RespondError(w http.ResponseWriter, data map[string]interface{}, log *Logger) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
	}
	log.Error(string(jsonString))
	w.WriteHeader(http.StatusBadRequest) //hide real error code
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hideError(data))
}

//Respond logs response and writes it into ResponseWriter
func Respond(w http.ResponseWriter, data map[string]interface{}, log *Logger) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info(string(jsonString))
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
