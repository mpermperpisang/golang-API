package biodata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-API/data"
	"github.com/golang-API/helper"
)

/*AddOneBiodata : add biodata user*/
func AddOneBiodata(w http.ResponseWriter, r *http.Request) {
	var newBio data.Field

	reqBody, err := ioutil.ReadAll(r.Body)
	helper.LogPanicln(err)

	if len(reqBody) > 0 {
		if err := json.Unmarshal(reqBody, &newBio); err != nil {
			helper.LogPanicln(err)
		}

		data.Datas = append(data.Datas, newBio)

		w.WriteHeader(http.StatusCreated)
		helper.PrettyJSON(w, newBio)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Request body cannot be empty")
	}
}
