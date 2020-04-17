package example

import "net/http"

/*Delete : contoh API delete*/
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "delete called",
		"owner": "mpermperpisang"
	}`))
}
