package helper

import (
	"encoding/json"
	"net/http"
)

const (
	empty = ""
	tab   = "\t"
)

/*PrettyJSON : tidy json*/
func PrettyJSON(w http.ResponseWriter, response interface{}) *json.Encoder {
	encoder := json.NewEncoder(w)

	encoder.SetIndent(empty, tab)

	err := encoder.Encode(response)
	LogPanicln(err)

	return encoder
}
