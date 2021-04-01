package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type IsError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func ResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
		return
	}
}

func ResponseJsonError(w http.ResponseWriter, data IsError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(data)
}

func GetDateTime() string {
	today := time.Now()

	day := today.Day()
	month := int(today.Month())
	year := today.Year()
	hour := today.Hour()
	minutes := today.Minute()
	seconds := today.Second()

	return fmt.Sprintf("%d-%d-%d T%d:%d:%d", day, month, year, hour, minutes, seconds)
}
