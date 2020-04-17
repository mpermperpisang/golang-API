package helper

import (
	"fmt"
	"net/http"
)

/*WriteStatusBodyText : variant status and body text*/
func WriteStatusBodyText(w http.ResponseWriter, statusCode int, text string) error {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, text)

	return nil
}
