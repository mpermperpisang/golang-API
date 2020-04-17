package biodata

import (
	"fmt"
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
	"github.com/gorilla/mux"
)

/*GetAllBiodata : get biodata user*/
func GetAllBiodata(w http.ResponseWriter, r *http.Request) {
	helper.PrettyJSON(w, data.Datas)
}

/*GetOneBiodata : get biodata of specific user*/
func GetOneBiodata(w http.ResponseWriter, r *http.Request) {
	bioID := mux.Vars(r)["id"]

	for _, singleBio := range data.Datas {
		if singleBio.ID == bioID {
			helper.PrettyJSON(w, singleBio)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "ID Not Found!")
		}
	}
}
