package example

import "net/http"

/*Patch : contoh API patch*/
func Patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "patch called",
		"owner": "mpermperpisang"
	}`))
}
