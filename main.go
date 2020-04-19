package main

import (
	"log"
	"net/http"

	"github.com/golang-API/endpoint"
	"github.com/golang-API/endpoint/biodata"
	"github.com/golang-API/endpoint/example"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	exampleEndpoint(r)
	biodataEndpoint(r)
	basicEndpoint(r)
	log.Fatal(http.ListenAndServe(":8181", r))
}

func exampleEndpoint(r *mux.Router) error {
	r.HandleFunc("/get", example.Get).Methods(http.MethodGet)
	r.HandleFunc("/post", example.Post).Methods(http.MethodPost)
	r.HandleFunc("/put", example.Put).Methods(http.MethodPut)
	r.HandleFunc("/patch", example.Patch).Methods(http.MethodPatch)
	r.HandleFunc("/delete", example.Delete).Methods(http.MethodDelete)

	return nil
}

func biodataEndpoint(r *mux.Router) error {
	// optional parameter
	r.Path("/biodata").Queries("id", "{id}").HandlerFunc(biodata.GetParamOneBiodata).Methods(http.MethodGet)
	r.Path("/biodata").Queries("id", "{id}").HandlerFunc(biodata.DeleteParamOneBiodata).Methods(http.MethodDelete)
	r.Path("/biodata").Queries("id", "{id}").HandlerFunc(biodata.UpdateParamOneBiodata).Methods(http.MethodPatch)

	// mandatory url
	r.HandleFunc("/biodata", biodata.GetAllBiodata).Methods(http.MethodGet)
	r.HandleFunc("/biodata/{id}", biodata.GetOneBiodata).Methods(http.MethodGet)
	r.HandleFunc("/biodata", biodata.AddOneBiodata).Methods(http.MethodPost)
	r.HandleFunc("/biodata/{id}", biodata.UpdateOneBiodata).Methods(http.MethodPatch)
	r.HandleFunc("/biodata/{id}", biodata.DeleteOneBiodata).Methods(http.MethodDelete)

	return nil
}

func basicEndpoint(r *mux.Router) error {
	r.HandleFunc("/", endpoint.NotFound).Methods(http.MethodGet)
	r.HandleFunc("/example", endpoint.Basic).Methods(http.MethodGet)

	return nil
}
