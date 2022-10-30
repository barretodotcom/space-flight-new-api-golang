package artc_http_func

import "net/http"

func Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Back-end Challenge 2022 🏅 - Space Flight News"))
}
