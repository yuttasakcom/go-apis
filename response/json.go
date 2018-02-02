package response

import (
	"encoding/json"
	"net/http"
)

var pretty = true

// JSON response with optional status code.
func JSON(w http.ResponseWriter, val interface{}, code ...int) {
	var b []byte
	var err error

	res := make(map[string]interface{})

	res["code"] = code[0]
	res["data"] = val

	if pretty {
		b, err = json.MarshalIndent(res, "", "  ")
	} else {
		b, err = json.Marshal(res)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	w.Write(b)
}

// Error responds with a generic status code response.
func Error(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
