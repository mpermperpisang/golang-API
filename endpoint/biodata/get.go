package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

func setGetResponseBody(w http.ResponseWriter, ID int) error {
	varInit()

	for _, singleBio := range data.Datas {
		if singleBio.ID == ID {
			responseBody = helper.PrettyJSON(w, singleBio)
		}
	}

	helper.ResponseBodyNotFound(responseBody, w)

	return nil
}

/*GetAllBiodata : get biodata user*/
func GetAllBiodata(w http.ResponseWriter, r *http.Request) {
	helper.PrettyJSON(w, data.Datas)
}

/*GetOneBiodata : get biodata of specific user*/
func GetOneBiodata(w http.ResponseWriter, r *http.Request) {
	varID(r)
	setGetResponseBody(w, bioID)
}

/*GetParamOneBiodata : get biodata of spesific user using param ID*/
func GetParamOneBiodata(w http.ResponseWriter, r *http.Request) {
	varParamBioID(r)
	setGetResponseBody(w, paramBioID)
}
