package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

func setPatchResponseBody(w http.ResponseWriter, r *http.Request, ID int) error {
	varInit()
	readBody(r)

	if len(reqBody) > 0 {
		unmarshallBody(reqBody, &bioField)

		for i, singleBio := range data.Datas {
			if singleBio.ID == ID {
				singleBio.Name = bioField.Name
				singleBio.Job = bioField.Job
				singleBio.Description = bioField.Description
				singleBio.Motto = bioField.Motto
				data.Datas = append(data.Datas[:i], singleBio)
				responseBody = helper.PrettyJSON(w, singleBio)
			}
		}
	} else {
		helper.WriteStatusBodyText(w, http.StatusBadRequest, helper.EmptyReqBody())
	}

	helper.ResponseBodyNotFound(responseBody, w)

	return nil
}

/*UpdateOneBiodata : update biodata user*/
func UpdateOneBiodata(w http.ResponseWriter, r *http.Request) {
	VarID(r)
	setPatchResponseBody(w, r, BioID)
}

/*UpdateParamOneBiodata : update biodata of spesific user using param ID*/
func UpdateParamOneBiodata(w http.ResponseWriter, r *http.Request) {
	VarParamBioID(r)
	setPatchResponseBody(w, r, ParamBioID)
}
