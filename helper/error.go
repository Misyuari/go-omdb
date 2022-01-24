package helper

import (
	"log"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func PanicHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	response := WrapResponse(http.StatusInternalServerError, err, "Internal Server Error")
	WriteToResponseBody(writer, response)
}
