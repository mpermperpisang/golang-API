package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

/*UpdateOneBiodata : update biodata user*/
func UpdateOneBiodata(w http.ResponseWriter, r *http.Request) {
	varBioID(r)
	readBody(r)

	responseBody = nil

	if len(reqBody) > 0 {
		unmarshallBody(reqBody, &bioField)

		for i, singleBio := range data.Datas {
			if singleBio.ID == bioID {
				singleBio.Name = bioField.Name
				singleBio.Job = bioField.Job
				singleBio.Description = bioField.Description
				singleBio.Motto = bioField.Motto
				data.Datas = append(data.Datas[:i], singleBio)
				responseBody = helper.PrettyJSON(w, singleBio)
			}
		}
	} else {
		helper.WriteStatusBodyText(w, http.StatusBadRequest, "Request body cannot be empty")
	}

	if responseBody == nil {
		helper.WriteStatusBodyText(w, http.StatusNotFound, "ID Not Found!")
	}
}
