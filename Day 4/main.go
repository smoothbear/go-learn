package main

import (
	"encoding/json"
	"net/http"
)

type userMap map[string]interface{}

func handlerFunc() http.Handler {
	mux := http.NewServeMux()
	user := userMap{}
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			v := []string{}
			for u := range user {
				v = append(v, u)
			}
			byteData, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(byteData)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	return mux
}
