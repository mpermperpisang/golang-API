package biodata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
	"github.com/gorilla/mux"
)

var responseBody *json.Encoder
var bioID int
var reqBody []byte
var bioField data.Field
var err error

func varInit() error {
	responseBody = nil

	return nil
}

func varID(r *http.Request) error {
	bioID, err = strconv.Atoi(mux.Vars(r)["id"])
	helper.LogPanicln(err)

	return nil
}

func readBody(r *http.Request) error {
	reqBody, err = ioutil.ReadAll(r.Body)
	helper.LogPanicln(err)

	return nil
}

func unmarshallBody(reqBody []byte, bio interface{}) error {
	err := json.Unmarshal(reqBody, bio)
	helper.LogPanicln(err)

	return nil
}
