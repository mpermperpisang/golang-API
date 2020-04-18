package biodata

import (
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

/*AddOneBiodata : add biodata user*/
func AddOneBiodata(w http.ResponseWriter, r *http.Request) {
	readBody(r)

	if len(reqBody) > 0 {
		unmarshallBody(reqBody, &bioField)

		data.Datas = append(data.Datas, bioField)

		w.WriteHeader(http.StatusCreated)
		helper.PrettyJSON(w, bioField)
	} else {
		helper.WriteStatusBodyText(w, http.StatusBadRequest, "Request body cannot be empty")
	}
}
