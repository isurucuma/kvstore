package routes

import (
	"encoding/json"
	"fmt"
	"kvstore/api/dto"
	"net/http"
)

func Routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /store", func(w http.ResponseWriter, r *http.Request) {
		// move this logic to a handler function
		// put the encoding decoding logincs to separate functions
		fmt.Println("PUT /store")
		max := http.MaxBytesReader(w, r.Body, 2) // handle this
		decoder := json.NewDecoder(max)
		tuple := &dto.Tuple{}
		decoder.Decode(tuple)

		w.WriteHeader(201)
		w.Write([]byte(fmt.Sprintf("%+v", tuple))) // send json as response
	})

	mux.HandleFunc("GET /store/{key}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /event")
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	})

	return mux
}
