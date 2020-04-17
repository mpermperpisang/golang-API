package helper

import (
	"encoding/json"
	"net/http"

	"github.com/golang-automation/features/helper"
)

const (
	empty = ""
	tab   = "\t"
)

/*PrettyJSON : tidy json*/
func PrettyJSON(w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)

	encoder.SetIndent(empty, tab)

	err := encoder.Encode(response)
	helper.LogPanicln(err)

	return nil
}
