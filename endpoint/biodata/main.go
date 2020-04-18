package biodata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
	"github.com/gorilla/mux"
)

var responseBody *json.Encoder
var bioID string
var reqBody []byte
var bioField data.Field

func varInit() error {
	responseBody = nil

	return nil
}

func varID(r *http.Request) error {
	bioID = mux.Vars(r)["id"]

	return nil
}

func readBody(r *http.Request) error {
	var err error

	reqBody, err = ioutil.ReadAll(r.Body)
	helper.LogPanicln(err)

	return nil
}

func unmarshallBody(reqBody []byte, bio interface{}) error {
	err := json.Unmarshal(reqBody, bio)
	helper.LogPanicln(err)

	return nil
}
