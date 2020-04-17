package example

import "net/http"

/*Get : contoh API get*/
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "get called",
		"owner": "mpermperpisang"
	}`))
}
