package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"rmock/models"
)

func ServeHttp(port int, jsonFilePath string) {
	fmt.Println("http called, server running on port", port, "serving", jsonFilePath)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	bytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}

	responses := models.Responses{}
	err = json.Unmarshal(bytes, &responses)
	if err != nil {
		panic(err)
	}

	for _, response := range responses.Responses {
		http.HandleFunc(response.Route, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != response.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				if _, err = w.Write([]byte("Method not allowed")); err != nil {
					panic(err)
				}
				return
			}

			bodyStr, err := json.Marshal(response.Body)
			if err != nil {
				panic(err)
			}
			_, err = w.Write(bodyStr)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(response.StatusCode)
		})
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
