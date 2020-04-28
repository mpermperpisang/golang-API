package endpoint

import "net/http"

/*NotFound : contoh API not found*/
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`
	{
		"message": "not found",
		"owner": "mpermperpisang",
		"command": [{
			"GET1": "/example",
			"GET2": "/biodata",
			"GET3": "/biodata/{id}",
			"POST": "/biodata",
			"PATCH": "/biodata/{id}",
			"DELETE": "/biodata/{id}"
		}]
	}`))
}
