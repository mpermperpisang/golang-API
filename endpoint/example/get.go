package example

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-API/endpoint/biodata"
)

/*Get : contoh API get*/
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "get called",
		"owner": "mpermperpisang"
	}`))
}

/*GetOneID : contoh API get with ID*/
func GetOneID(w http.ResponseWriter, r *http.Request) {
	biodata.VarID(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"id": ` + strconv.Itoa(biodata.BioID) + `,
		"owner": "mpermperpisang"
	}`))
}

/*GetParamOneID : contoh API get with optional param ID*/
func GetParamOneID(w http.ResponseWriter, r *http.Request) {
	biodata.VarParamBioID(r)
	fmt.Println(biodata.ParamBioID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"id": "` + strconv.Itoa(biodata.ParamBioID) + `",
		"owner": "mpermperpisang"
	}`))
}
