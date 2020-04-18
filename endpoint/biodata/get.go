package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

/*GetAllBiodata : get biodata user*/
func GetAllBiodata(w http.ResponseWriter, r *http.Request) {
	helper.PrettyJSON(w, data.Datas)
}

/*GetOneBiodata : get biodata of specific user*/
func GetOneBiodata(w http.ResponseWriter, r *http.Request) {
	varInit()
	varID(r)

	for _, singleBio := range data.Datas {
		if singleBio.ID == bioID {
			responseBody = helper.PrettyJSON(w, singleBio)
		}
	}

	if responseBody == nil {
		helper.WriteStatusBodyText(w, http.StatusNotFound, helper.NotFoundID())
	}
}
