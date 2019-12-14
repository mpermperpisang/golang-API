package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "get called",
		"owner": "mpermperpisang"
	}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{
		"message": "post called",
		"owner": "mpermperpisang"
	}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{
		"message": "put called",
		"owner": "mpermperpisang"
	}`))
}

func patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "patch called",
		"owner": "mpermperpisang"
	}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "delete called",
		"owner": "mpermperpisang"
	}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`
	{
		"message": "not found",
		"owner": "mpermperpisang",
		"command": [
			{
				"GET": "/get",
				"POST": "/post",
				"PUT": "/put",
				"PATCH": "/patch",
				"DELETE": "/delete",
			}
		]
	}`))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get", get).Methods(http.MethodGet)
	r.HandleFunc("/post", post).Methods(http.MethodPost)
	r.HandleFunc("/put", put).Methods(http.MethodPut)
	r.HandleFunc("/patch", patch).Methods(http.MethodPatch)
	r.HandleFunc("/delete", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8181", r))
}
