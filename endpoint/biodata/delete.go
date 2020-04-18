package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

/*DeleteOneBiodata : delete biodata user*/
func DeleteOneBiodata(w http.ResponseWriter, r *http.Request) {
	varBioID(r)

	responseBody = nil

	for i, singleBio := range data.Datas {
		if singleBio.ID == bioID {
			data.Datas = append(data.Datas[:i], data.Datas[i+1:]...)
			responseBody = helper.PrettyJSON(w, singleBio)
		}
	}

	if responseBody == nil {
		helper.WriteStatusBodyText(w, http.StatusNotFound, "ID Not Found!")
	}
}
