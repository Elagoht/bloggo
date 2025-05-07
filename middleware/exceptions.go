package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Elagoht/bloggo/utils"
)

type HandlerFuncWithError func(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError

func Handle(function HandlerFuncWithError) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := function(writer, request); err != nil {

			status := err.Status
			if status == 0 { // Created with zero value
				status = http.StatusInternalServerError
			}

			response := map[string]string{"error": err.Message}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(status)
			json.NewEncoder(writer).Encode(response)

			if err.Err != nil {
				log.Println("Internal server error:", err.Err.Error())
			}
		}
	}
}
