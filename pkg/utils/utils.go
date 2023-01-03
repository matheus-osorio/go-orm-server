package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, expectedModel interface{}) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), expectedModel)

	if err != nil {
		panic(err)
	}
}

func SetDefaultHeaders(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set("Content-Type", "aplication/json")
	return writer
}
