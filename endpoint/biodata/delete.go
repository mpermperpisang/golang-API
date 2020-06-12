package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

func setDeleteResponseBody(w http.ResponseWriter, ID int) error {
	varInit()

	for i, singleBio := range data.Datas {
		if singleBio.ID == ID {
			data.Datas = append(data.Datas[:i], data.Datas[i+1:]...)
			responseBody = helper.PrettyJSON(w, singleBio)
		}
	}

	helper.ResponseBodyNotFound(responseBody, w)

	return nil
}

/*DeleteOneBiodata : delete biodata user*/
func DeleteOneBiodata(w http.ResponseWriter, r *http.Request) {
	VarID(r)
	setDeleteResponseBody(w, BioID)
}

/*DeleteParamOneBiodata : delete biodata of spesific user using param ID*/
func DeleteParamOneBiodata(w http.ResponseWriter, r *http.Request) {
	VarParamBioID(r)
	setDeleteResponseBody(w, ParamBioID)
}
