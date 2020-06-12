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
var BioID, ParamBioID int
var reqBody []byte
var bioField data.Field
var err error

func varInit() error {
	responseBody = nil

	return nil
}

/*VarID : content ID in url*/
func VarID(r *http.Request) error {
	BioID, err = strconv.Atoi(mux.Vars(r)["id"])
	helper.LogPanicln(err)

	return nil
}

/*VarParamBioID : content ID in param url*/
func VarParamBioID(r *http.Request) error {
	ParamBioID, err = strconv.Atoi(r.FormValue("id"))
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
