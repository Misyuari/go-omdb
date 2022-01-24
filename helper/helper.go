package helper

import (
	"encoding/json"
	"net/http"
)

func RequestGet(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ConvertGrpcStruct(from interface{}, to interface{}) error {
	bytes, _ := json.Marshal(from)
	return json.Unmarshal(bytes, to)
}
