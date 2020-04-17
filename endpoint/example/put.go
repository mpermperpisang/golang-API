package example

import "net/http"

/*Put : contoh API put*/
func Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{
		"message": "put called",
		"owner": "mpermperpisang"
	}`))
}
