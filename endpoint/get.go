package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
	"github.com/gorilla/mux"
)

const (
	empty = ""
	tab   = "\t"
)

/*GetAllBiodata : get biodata user*/
func GetAllBiodata(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data.Datas)
	if err != nil {
		helper.LogPanicln(err)
	}
}

/*GetOneBiodata : get biodata of specific user*/
func GetOneBiodata(w http.ResponseWriter, r *http.Request) {
	bioID := mux.Vars(r)["id"]
	encoder := json.NewEncoder(w)

	encoder.SetIndent(empty, tab)

	for _, singleBio := range data.Datas {
		if singleBio.ID == bioID {
			err := encoder.Encode(singleBio)
			if err != nil {
				helper.LogPanicln(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "ID Not Found!")
		}
	}
}
