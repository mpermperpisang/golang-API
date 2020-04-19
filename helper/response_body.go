package helper

import (
	"encoding/json"
	"net/http"
)

/*ResponseBodyNotFound : check if response body is null then status must be not found*/
func ResponseBodyNotFound(response *json.Encoder, w http.ResponseWriter) {
	if response == nil {
		WriteStatusBodyText(w, http.StatusNotFound, NotFoundID())
	}
}
