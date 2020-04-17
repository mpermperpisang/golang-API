package example

import "net/http"

/*Post : contoh API post*/
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{
		"message": "post called",
		"owner": "mpermperpisang"
	}`))
}
