package endpoint

import "net/http"

/*Basic : contoh endpoint API*/
func Basic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	{
		"message": "basic example",
		"owner": "mpermperpisang",
		"command": [
			{
				"GET": "/get",
				"POST": "/post",
				"PUT": "/put",
				"PATCH": "/patch",
				"DELETE": "/delete"
			}
		]
	}`))
}
